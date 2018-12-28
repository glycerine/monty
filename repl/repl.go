// Package repl provides a read/eval/print loop for Starlark.
//
// It supports readline-style command editing,
// and interrupts through Control-C.
//
// If an input line can be parsed as an expression,
// the REPL parses and evaluates it and prints its result.
// Otherwise the REPL reads lines until a blank line,
// then tries again to parse the multi-line input as an
// expression. If the input still cannot be parsed as an expression,
// the REPL parses and executes it as a file (a list of statements),
// for side effects.
package repl

// TODO(adonovan):
//
// - Unparenthesized tuples are not parsed as a single expression:
//     >>> (1, 2)
//     (1, 2)
//     >>> 1, 2
//     ...
//     >>>
//   This is not necessarily a bug.

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/chzyer/readline"
	"github.com/glycerine/monty/starlark"
	"github.com/glycerine/monty/syntax"
)

var interrupted = make(chan os.Signal, 1)

// REPL executes a read, eval, print loop.
//
// Before evaluating each expression, it sets the Starlark thread local
// variable named "context" to a context.Context that is cancelled by a
// SIGINT (Control-C). Client-supplied global functions may use this
// context to make long-running operations interruptable.
//
func REPL(thread *starlark.Thread, globals *starlark.StringDict) {
	signal.Notify(interrupted, os.Interrupt)
	defer signal.Stop(interrupted)

	rl, err := readline.New(">>> ")
	if err != nil {
		PrintError(err)
		return
	}
	prep := withPrepend(rl)
	defer rl.Close()
	for {
		if err := rep(prep, thread, globals); err != nil {
			if err == readline.ErrInterrupt {
				fmt.Println(err)
				continue
			}
			break
		}
	}
	fmt.Println()
}

var newline = []byte{'\n'}

// rep reads, evaluates, and prints one item.
//
// It returns an error (possibly readline.ErrInterrupt)
// only if readline failed. Starlark errors are printed.
func rep(rl *prepender, thread *starlark.Thread, globals *starlark.StringDict) error {
	rl.reset()
	// Each item gets its own context,
	// which is cancelled by a SIGINT.
	//
	// Note: during Readline calls, Control-C causes Readline to return
	// ErrInterrupt but does not generate a SIGINT.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		select {
		case <-interrupted:
			cancel()
		case <-ctx.Done():
		}
	}()

	thread.SetLocal("context", ctx)

	rl.SetPrompt(">>> ")

	line, err := rl.ReadSlice()
	if err != nil {
		return err // may be ErrInterrupt
	}

	if l := bytes.TrimSpace(line); len(l) == 0 || l[0] == '#' {
		return nil // blank or comment
	}

	rl.prepend(line)

	// If the line contains a well-formed expression, evaluate it.
	x, err := syntax.ParseExpr("<stdin>", rl, 0)
	if err == nil {
		if v, err := starlark.EvalExpr(thread, "<stdin>", x, globals); err != nil {
			PrintError(err)
		} else if v != starlark.None {
			fmt.Println(v)
		}
		return nil
	}

	rl.startRecord()
	rl.prepend(line)

	// If the input so far is a single load or assignment statement,
	// execute it without waiting for a blank line.
	if f, err := syntax.Parse("<stdin>", rl, 0); err == nil && f != nil && len(f.Stmts) == 1 {
		switch f.Stmts[0].(type) {
		case *syntax.AssignStmt, *syntax.LoadStmt:
			// Execute it as a file.
			rl.restore()
			oneStmt := rl.prefix
			if err := execFileNoFreeze(thread, oneStmt, globals); err != nil {
				PrintError(err)
			}
			return nil
		}
	}

	rl.restore()
	rl.stopRecording()

	// Otherwise assume it is the first of several
	// comprising a file, followed by a blank line.
	var buf bytes.Buffer
	buf.Write(rl.prefix)
	buf.Write(newline)
	for {
		rl.SetPrompt("... ")
		line, err := rl.wrapped.ReadSlice()
		if err != nil {
			return err // may be ErrInterrupt
		}
		if l := bytes.TrimSpace(line); len(l) == 0 {
			break // blank
		}
		buf.Write(line)
		buf.Write(newline)
	}
	text := buf.Bytes()

	// Try parsing it once more as an expression,
	// such as a call spread over several lines:
	//   f(
	//     1,
	//     2
	//   )
	if _, err := syntax.ParseExpr("<stdin>", text, 0); err == nil {
		if v, err := starlark.Eval(thread, "<stdin>", text, globals); err != nil {
			PrintError(err)
		} else if v != starlark.None {
			fmt.Println(v)
		}
		return nil
	}

	// Execute it as a file.
	if err := execFileNoFreeze(thread, text, globals); err != nil {
		PrintError(err)
	}

	return nil
}

// execFileNoFreeze is starlark.ExecFile without globals.Freeze().
func execFileNoFreeze(thread *starlark.Thread, src interface{}, globals *starlark.StringDict) error {
	_, prog, err := starlark.SourceProgram("<stdin>", src, globals.Has)
	if err != nil {
		return err
	}

	res, err := prog.Init(thread, globals)

	// The global names from the previous call become
	// the predeclared names of this call.

	// Copy globals back to the caller's map.
	// If execution failed, some globals may be undefined.
	for k, v := range res.Map {
		globals.Map[k] = v
	}

	return err
}

// PrintError prints the error to stderr,
// or its backtrace if it is a Starlark evaluation error.
func PrintError(err error) {
	if evalErr, ok := err.(*starlark.EvalError); ok {
		fmt.Fprintln(os.Stderr, evalErr.Backtrace())
	} else {
		fmt.Fprintln(os.Stderr, err)
	}
}

// MakeLoad returns a simple sequential implementation of module loading
// suitable for use in the REPL.
// Each function returned by MakeLoad accesses a distinct private cache.
func MakeLoad() func(thread *starlark.Thread, module string) (*starlark.StringDict, error) {
	type entry struct {
		globals *starlark.StringDict
		err     error
	}

	var cache = make(map[string]*entry)

	return func(thread *starlark.Thread, module string) (*starlark.StringDict, error) {
		e, ok := cache[module]
		if e == nil {
			if ok {
				// request for package whose loading is in progress
				return nil, fmt.Errorf("cycle in load graph")
			}

			// Add a placeholder to indicate "load in progress".
			cache[module] = nil

			// Load it.
			thread := &starlark.Thread{Name: "exec " + module, Load: thread.Load}
			globals, err := starlark.ExecFile(thread, module, nil, nil)
			e = &entry{globals, err}

			// Update the cache.
			cache[module] = e
		}
		return e.globals, e.err
	}
}

// prepender allows us to use the readline library
// to handle multiline strings, even when we attempt
// multiple parses of the same line.
type prepender struct {
	wrapped   *readline.Instance
	prefix    []byte
	recording bool
	record    []byte
}

func withPrepend(rl *readline.Instance) (p *prepender) {
	return &prepender{
		wrapped: rl,
	}
}

func (p *prepender) SetPrompt(s string) {
	p.wrapped.SetPrompt(s)
}

func (p *prepender) prepend(line []byte) {
	p.prefix = line
}

func (p *prepender) clearPrefix() {
	p.prefix = p.prefix[:0]
}

func (p *prepender) reset() {
	p.record = p.record[:0]
	p.prefix = p.prefix[:0]
	p.recording = false
}

func (p *prepender) startRecord() {
	p.recording = true
	if p.record == nil {
		p.record = make([]byte, 0, 64)
	} else {
		p.record = p.record[:0]
	}
}

func (p *prepender) stopRecording() {
	p.recording = false
}

func (p *prepender) restore() {
	p.prefix = p.record
	p.record = p.record[:0]
	p.recording = false
}

func (p *prepender) ReadSlice() (by []byte, err error) {
	if len(p.prefix) != 0 {
		by = p.prefix
		p.prefix = p.prefix[:0]

	} else {
		by, err = p.wrapped.ReadSlice()
		if err == nil {
			by = append(by, '\n') // restore the byte that readline trimmed off.
		}
	}
	if p.recording {
		p.record = append(p.record, by...)
	}
	return
}
