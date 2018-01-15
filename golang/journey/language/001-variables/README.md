# Variables in Golang

## Pre-requisites
See the [Setting up Go in README.md](../../README.md#setting-up-go) to get started.

## What's Covered
Variable assignments and different ways of doing it. See the source code at [`./variables.go`](./variables.go).

Covered:

- variable assignment
- variable types
- variable passing

Build it with:

```bash
go build variables.go
```

You should now be able to run it with:

```bash
./variables.go
```

## Learnings

- Golang has type inference despite being a strongly typed language
- Golang has pointers!!! Use them as in C/C++
- Strings must be in double quotes
  - *Try changing `"\n"` to `'\n'`*
- Formatting of `fmt.Printf`
  - See [https://golang.org/pkg/fmt/](https://golang.org/pkg/fmt/) for more!
- All variables are immutable in Golang, see `f := a` example
- Syntatic sugars are available
  - Use `:=` for single line multiple variable assignments