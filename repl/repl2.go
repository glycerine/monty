package repl

import (
	"fmt"
	"os"

	"github.com/glycerine/monty/starlarkstruct"

	"github.com/starlight-go/starlight"
	"github.com/starlight-go/starlight/convert"

	"github.com/glycerine/monty/starlark"
	"github.com/glycerine/monty/verb"

	shadow_bytes "github.com/glycerine/monty/lib/bytes"
	shadow_encoding "github.com/glycerine/monty/lib/encoding"
	shadow_encoding_binary "github.com/glycerine/monty/lib/encoding/binary"
	shadow_errors "github.com/glycerine/monty/lib/errors"
	shadow_fmt "github.com/glycerine/monty/lib/fmt"
	shadow_io "github.com/glycerine/monty/lib/io"
	shadow_io_ioutil "github.com/glycerine/monty/lib/io/ioutil"
	shadow_math "github.com/glycerine/monty/lib/math"
	shadow_math_rand "github.com/glycerine/monty/lib/math/rand"
	shadow_os "github.com/glycerine/monty/lib/os"
	shadow_reflect "github.com/glycerine/monty/lib/reflect"
	shadow_regexp "github.com/glycerine/monty/lib/regexp"
	shadow_runtime "github.com/glycerine/monty/lib/runtime"
	shadow_runtime_debug "github.com/glycerine/monty/lib/runtime/debug"
	shadow_strconv "github.com/glycerine/monty/lib/strconv"
	shadow_strings "github.com/glycerine/monty/lib/strings"
	shadow_sync "github.com/glycerine/monty/lib/sync"
	shadow_sync_atomic "github.com/glycerine/monty/lib/sync/atomic"
	shadow_time "github.com/glycerine/monty/lib/time"

	shadow_archive_tar "github.com/glycerine/monty/lib/archive/tar"
	shadow_archive_zip "github.com/glycerine/monty/lib/archive/zip"
	shadow_bufio "github.com/glycerine/monty/lib/bufio"
	shadow_compress_bzip2 "github.com/glycerine/monty/lib/compress/bzip2"
	shadow_compress_flate "github.com/glycerine/monty/lib/compress/flate"
	shadow_compress_gzip "github.com/glycerine/monty/lib/compress/gzip"
	shadow_compress_lzw "github.com/glycerine/monty/lib/compress/lzw"
	shadow_compress_zlib "github.com/glycerine/monty/lib/compress/zlib"
	shadow_context "github.com/glycerine/monty/lib/context"
	shadow_database_sql "github.com/glycerine/monty/lib/database/sql"
	shadow_database_sql_driver "github.com/glycerine/monty/lib/database/sql/driver"
	shadow_encoding_base64 "github.com/glycerine/monty/lib/encoding/base64"
	shadow_encoding_csv "github.com/glycerine/monty/lib/encoding/csv"
	shadow_encoding_json "github.com/glycerine/monty/lib/encoding/json"
	shadow_encoding_xml "github.com/glycerine/monty/lib/encoding/xml"
	shadow_flag "github.com/glycerine/monty/lib/flag"
	shadow_html "github.com/glycerine/monty/lib/html"
	shadow_html_template "github.com/glycerine/monty/lib/html/template"
	shadow_math_big "github.com/glycerine/monty/lib/math/big"
	shadow_math_bits "github.com/glycerine/monty/lib/math/bits"
	shadow_math_cmplx "github.com/glycerine/monty/lib/math/cmplx"
	shadow_net "github.com/glycerine/monty/lib/net"
	shadow_net_http "github.com/glycerine/monty/lib/net/http"
	shadow_os_exec "github.com/glycerine/monty/lib/os/exec"
	shadow_os_signal "github.com/glycerine/monty/lib/os/signal"
	shadow_path "github.com/glycerine/monty/lib/path"
	shadow_path_filepath "github.com/glycerine/monty/lib/path/filepath"
)

var vv = verb.VV

type MontyEnv struct {
	InitDone bool

	// GoGlobal -> translated into GlobalDict
	GoGlobal   map[string]interface{}
	GlobalDict *starlark.StringDict

	ScriptCache *starlight.Cache
	Thread      *starlark.Thread
}

func NewMontyEnv() *MontyEnv {
	e := &MontyEnv{
		GoGlobal: make(map[string]interface{}),
	}
	e.Init()
	return e
}

func (env *MontyEnv) Init() {
	if env.InitDone {
		return
	}

	env.GoGlobal["strconv"] = shadow_strconv.Pkg
	env.GoGlobal["strings"] = shadow_strings.Pkg
	env.GoGlobal["regexp"] = shadow_regexp.Pkg
	env.GoGlobal["bytes"] = shadow_bytes.Pkg
	env.GoGlobal["fmt"] = shadow_fmt.Pkg
	env.GoGlobal["reflect"] = shadow_reflect.Pkg
	env.GoGlobal["ioutil"] = shadow_io_ioutil.Pkg
	env.GoGlobal["io"] = shadow_io.Pkg
	env.GoGlobal["runtime"] = shadow_runtime.Pkg
	env.GoGlobal["debug"] = shadow_runtime_debug.Pkg
	env.GoGlobal["binary"] = shadow_encoding_binary.Pkg
	env.GoGlobal["encoding"] = shadow_encoding.Pkg
	env.GoGlobal["math"] = shadow_math.Pkg
	env.GoGlobal["rand"] = shadow_math_rand.Pkg
	env.GoGlobal["time"] = shadow_time.Pkg
	env.GoGlobal["os"] = shadow_os.Pkg
	env.GoGlobal["errors"] = shadow_errors.Pkg
	env.GoGlobal["sync"] = shadow_sync.Pkg
	env.GoGlobal["atomic"] = shadow_sync_atomic.Pkg

	env.GoGlobal["net"] = shadow_net.Pkg
	env.GoGlobal["http"] = shadow_net_http.Pkg
	env.GoGlobal["path"] = shadow_path.Pkg
	env.GoGlobal["filepath"] = shadow_path_filepath.Pkg
	env.GoGlobal["exec"] = shadow_os_exec.Pkg
	env.GoGlobal["signal"] = shadow_os_signal.Pkg
	env.GoGlobal["tar"] = shadow_archive_tar.Pkg
	env.GoGlobal["zip"] = shadow_archive_zip.Pkg
	env.GoGlobal["bufio"] = shadow_bufio.Pkg
	env.GoGlobal["gzip"] = shadow_compress_gzip.Pkg
	env.GoGlobal["bzip2"] = shadow_compress_bzip2.Pkg
	env.GoGlobal["zlib"] = shadow_compress_zlib.Pkg
	env.GoGlobal["zlib"] = shadow_compress_zlib.Pkg
	env.GoGlobal["lzw"] = shadow_compress_lzw.Pkg
	env.GoGlobal["flate"] = shadow_compress_flate.Pkg
	env.GoGlobal["context"] = shadow_context.Pkg
	env.GoGlobal["sql"] = shadow_database_sql.Pkg
	env.GoGlobal["driver"] = shadow_database_sql_driver.Pkg
	env.GoGlobal["json"] = shadow_encoding_json.Pkg
	env.GoGlobal["base64"] = shadow_encoding_base64.Pkg
	env.GoGlobal["csv"] = shadow_encoding_csv.Pkg
	env.GoGlobal["xml"] = shadow_encoding_xml.Pkg
	env.GoGlobal["flag"] = shadow_flag.Pkg
	env.GoGlobal["html"] = shadow_html.Pkg
	env.GoGlobal["template"] = shadow_html_template.Pkg
	env.GoGlobal["big"] = shadow_math_big.Pkg
	env.GoGlobal["bits"] = shadow_math_bits.Pkg
	env.GoGlobal["cmplx"] = shadow_math_cmplx.Pkg

	env.GoGlobal["string"] = starlark.GenericAsString

	// get someplace. TODO: improve where we look for scripts.
	dir, err := os.Getwd()
	panicOn(err)
	scriptCache := starlight.New(dir)

	var dict *starlark.StringDict
	if len(env.GoGlobal) > 0 {
		dict, err = convert.MakeStringDict(env.GoGlobal)
		panicOn(err)
	} else {
		dict = starlark.NewStringDict(0)
	}

	// freeze them all, to avoid user overwriting them.
	for nm, pkg := range dict.Map {
		pkg.Freeze()
		// name the built in packages
		sd, ok := pkg.(*starlark.StringDict)
		if ok {
			sd.PackageName = nm
		}
	}

	// add the struct constructor.
	dict.Map["struct"] = starlark.NewBuiltin("struct", starlarkstruct.Make)

	// add the package constructor.
	dict.Map["package"] = starlark.NewBuiltin("package", PackageMaker)

	env.GlobalDict = dict
	env.ScriptCache = scriptCache
	env.InitDone = true

	env.Thread = &starlark.Thread{
		Name: "montyREPL",
		Load: env.ScriptCache.Load,
	}

}

func (env *MontyEnv) Eval(code string) error {

	if !env.InitDone {
		env.Init()
	}

	//vv("pre-global='%#v'", env.GlobalDict.String())

	var err error
	var back *starlark.StringDict

	back, err = starlark.ExecFile(env.Thread, "eval.sky", code, env.GlobalDict)

	//vv("back='%#v'", back)
	//vv("global='%#v'", env.GlobalDict.String())

	// merge back into globals
	for k, v := range back.Map {
		env.GlobalDict.Map[k] = v
	}

	if err != nil {
		return fmt.Errorf("monty evaluation error: '%v'; on code: '%s'", err, string(code))
	}
	return nil
}

func panicOn(err error) {
	if err != nil {
		panic(err)
	}
}

func PackageMaker(_ *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	d := starlark.NewStringDict(0)
	if len(kwargs) == 0 && len(args) == 0 {
		return d, nil
	}
	if len(args) > 0 {
		// name is first arg
		s, ok := args[0].(starlark.String)
		if ok {
			d.PackageName = string(s)
		}
	} else {
		// search for "name"
		if len(kwargs) > 0 {
			for _, tup := range kwargs {
				k, ok := tup[0].(starlark.String)
				if ok && string(k) == "name" {
					v, ok := tup[1].(starlark.String)
					if ok {
						d.PackageName = string(v)
					}
				}
			}
		}
	}
	return d, nil
}
