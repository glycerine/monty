package repl

import (
	"fmt"
	"os"

	"github.com/starlight-go/starlight"
	"github.com/starlight-go/starlight/convert"

	"github.com/glycerine/monty/starlark"
	"github.com/glycerine/monty/verb"
)

var vv = verb.VV

type MontyEnv struct {
	InitDone bool

	// GoGlobal are translated into GlobalDict
	GoGlobal   map[string]interface{}
	GlobalDict starlark.StringDict

	ScriptCache *starlight.Cache
	Thread      *starlark.Thread
}

func (env *MontyEnv) Init() {
	if env.InitDone {
		return
	}
	// get someplace. TODO: improve where we look for scripts.
	dir, err := os.Getwd()
	panicOn(err)
	scriptCache := starlight.New(dir)

	var dict starlark.StringDict
	if len(env.GoGlobal) > 0 {
		dict, err = convert.MakeStringDict(env.GoGlobal)
		panicOn(err)
	} else {
		dict = make(map[string]starlark.Value)
	}

	env.GlobalDict = dict
	env.ScriptCache = scriptCache
	env.InitDone = true

	env.Thread = &starlark.Thread{
		Name: "zlispREPL",
		Load: env.ScriptCache.Load,
	}

}

func (env *MontyEnv) Eval(code string) error {

	if !env.InitDone {
		env.Init()
	}

	//vv("pre-global='%#v'", env.GlobalDict.String())

	var err error
	var back starlark.StringDict

	back, err = starlark.ExecFile(env.Thread, "eval.sky", code, env.GlobalDict)

	//vv("back='%#v'", back)
	//vv("global='%#v'", env.GlobalDict.String())

	// merge back into globals
	for k, v := range back {
		env.GlobalDict[k] = v
	}

	if err != nil {
		return fmt.Errorf("starlark evaluation error: '%v'; on code: '%s'", err, string(code))
	}
	return nil
}

func panicOn(err error) {
	if err != nil {
		panic(err)
	}
}
