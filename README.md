# Monty

Monty is a dialect of Python, supplemented with the Go standard
library. Calls into compiled library functions leverage that highly
tuned and well tested code. Monty provides a subset of Python.
The Monty interpreter is pure Go, calls compiled Go by reflection,
and runs without any CGO dependencies.

It is licensed under a 3-clause BSD style license.

Monty is a fork of the starlark-go project.
That code is available at https://go.starlark.net ,
also known as https://github.com/google/starlark-go.

Monty uses code from Nate Finch's https://github.com/starlight-go/starlight
project, which is licensed under the MIT license.

This talk by Alan Donovan, https://www.youtube.com/watch?v=9P_YKVhncWI, is a good
introduction to the original motivation for Starlark (nee Skylark) as a configuration
language for the Bazel build system.

An adaptation of the README for Starlark-Go follows.

## Starlark in Go

Starlark in Go is an interpreter for Starlark, implemented in Go.
Starlark was formerly known as Skylark.

Starlark is a dialect of Python intended for use as a configuration language.
Like Python, it is an untyped dynamic language with high-level data
types, first-class functions with lexical scope, and garbage collection.
Unlike CPython, independent Starlark threads execute in parallel, so
Starlark workloads scale well on parallel machines.
Starlark is a small and simple language with a familiar and highly
readable syntax. You can use it as an expressive notation for
structured data, defining functions to eliminate repetition, or you
can use it to add scripting capabilities to an existing application.

A Starlark interpreter is typically embedded within a larger
application, and the application may define additional domain-specific
functions and data types beyond those provided by the core language.
For example, Starlark was originally developed for the
[Bazel build tool](https://bazel.build).
Bazel uses Starlark as the notation both for its BUILD files (like
Makefiles, these declare the executables, libraries, and tests in a
directory) and for [its macro
language](https://docs.bazel.build/versions/master/skylark/language.html),
through which Bazel is extended with custom logic to support new
languages and compilers.


## Documentation

* Language definition: [doc/spec.md](doc/spec.md)

* About the Go implementation: [doc/impl.md](doc/impl.md)

* API documentation: [godoc.org/go.starlark.net/starlark](https://godoc.org/go.starlark.net/starlark)

* Mailing list: [starlark-go](https://groups.google.com/forum/#!forum/starlark-go)

* Issue tracker: [https://github.com/google/starlark-go/issues](https://github.com/google/starlark-go/issues)

* Travis CI: [![Travis CI](https://travis-ci.org/google/starlark-go.svg) https://travis-ci.org/google/starlark-go](https://travis-ci.org/google/starlark-go)

### Getting started

Build the code:

```shell
# check out the code and dependencies,
# and install interpreter in $GOPATH/bin
$ go get -u github.com/glycerine/monty/cmd/monty
```

Run the interpreter:

```
$ cat coins.star
coins = {
  'dime': 10,
  'nickel': 5,
  'penny': 1,
  'quarter': 25,
}
print('By name:\t' + ', '.join(sorted(coins.keys())))
print('By value:\t' + ', '.join(sorted(coins.keys(), key=coins.get)))

$ monty coins.star
By name:	dime, nickel, penny, quarter
By value:	penny, nickel, dime, quarter
```

Interact with the read-eval-print loop (REPL):

```
$ monty
>>> def fibonacci(n):
...    res = list(range(n))
...    for i in res[2:]:
...        res[i] = res[i-2] + res[i-1]
...    return res
...
>>> fibonacci(10)
[0, 1, 1, 2, 3, 5, 8, 13, 21, 34]
>>>
```

When you have finished, type `Ctrl-D` to close the REPL's input stream. 


### Credits

Starlark was designed and implemented in Java by
Ulf Adams,
Lukács Berki,
Jon Brandvein,
John Field,
Laurent Le Brun,
Dmitry Lomov,
Damien Martin-Guillerez,
Vladimir Moskva, and
Florian Weikert,
standing on the shoulders of the Python community.
The Go implementation was written by Alan Donovan and Jay Conrod;
its scanner was derived from one written by Russ Cox.

### Legal

Monty is Copyright(c) 2018 Jason E. Aten and the Monty authors.
All rights reserved.

Starlark in Go is Copyright (c) 2018 The Bazel Authors.
All rights reserved.

It is provided under a 3-clause BSD license:
[LICENSE](https://github.com/glycerine/monty/blob/master/LICENSE).

Starlark in Go is not an official Google product.
