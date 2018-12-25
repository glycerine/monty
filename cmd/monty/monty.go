// Copyright 2017 The Monty Authors. All rights reserved.
// Copyright 2017 The Bazel Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The monty command interprets a Monty file.
// With no arguments, it starts a read-eval-print loop (REPL).
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"sort"
	"strings"

	"github.com/glycerine/monty/repl"
	"github.com/glycerine/monty/resolve"
	"github.com/glycerine/monty/starlark"

	"github.com/glycerine/monty/starlarktest"
	"github.com/glycerine/monty/verb"
)

var vv = verb.VV

// flags
var (
	cpuprofile = flag.String("cpuprofile", "", "gather CPU profile in this file")
	showenv    = flag.Bool("showenv", false, "on success, print final global environment")
	execprog   = flag.String("c", "", "execute program `prog`")
)

func init() {
	resolve.AllowNestedDef = true      // allow def statements within function bodies
	resolve.AllowLambda = true         // allow lambda expressions
	resolve.AllowFloat = true          // allow floating point literals, the 'float' built-in, and x / y
	resolve.AllowSet = true            // allow the 'set' built-in
	resolve.AllowBitwise = true        // allow bitwise operations
	resolve.AllowRecursion = true      // allow while statements and recursive functions
	resolve.AllowGlobalReassign = true // allow reassignment of globals, and if/for/while statements at top level
	//resolve.AllowAddressable = true
}

func main() {
	log.SetPrefix("monty: ")
	log.SetFlags(0)
	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal(err)
		}
		defer pprof.StopCPUProfile()
	}

	thread := &starlark.Thread{Load: repl.MakeLoad()}
	//globals := make(starlark.StringDict)

	globals, err := starlarktest.LoadAssertModule()
	if err != nil {
		panic(err)
	}

	switch {
	case flag.NArg() == 1 || *execprog != "":
		var (
			filename string
			src      interface{}
			err      error
		)
		if *execprog != "" {
			// Execute provided program.
			filename = "cmdline"
			src = *execprog
		} else {
			// Execute specified file.
			filename = flag.Arg(0)
		}
		thread.Name = "exec " + filename
		back, err := starlark.ExecFile(thread, filename, src, globals)
		if err != nil {
			repl.PrintError(err)
			os.Exit(1)
		}

		vv("before merge, globals = '%s'", globals)
		vv("back = '%s'", back)

		// merge back into globals
		for k, v := range back {
			globals[k] = v
		}

		vv("after merge globals = '%s'", globals)

	case flag.NArg() == 0:
		fmt.Println("Welcome to Monty (github.com/glycerine/monty)")
		thread.Name = "REPL"

		vv("before REPL, globals = '%s'", globals)

		repl.REPL(thread, globals)
	default:
		log.Fatal("want at most one Monty file name")
	}

	// Print the global environment.
	if *showenv {
		var names []string
		for name := range globals {
			if !strings.HasPrefix(name, "_") {
				names = append(names, name)
			}
		}
		sort.Strings(names)
		for _, name := range names {
			fmt.Fprintf(os.Stderr, "%s = %s\n", name, globals[name])
		}
	}
}
