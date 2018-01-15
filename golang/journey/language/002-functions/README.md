# Variables in Golang

## Pre-requisites
See the [Setting up Go in README.md](../../README.md#setting-up-go) to get started.

## What's Covered
Functions and function calls. See the source code at [`./functions.go`](./functions.go).

Covered are:

- defining a function
- calling a function from the same package
- calling a function from a different package
- argument definition
- variable argument list (variadic functions)
- closures
- defer

Build and run it with:

```bash
go run *.go
```

## Learnings

- Avoid underscore_case, use camelCase for function names
- Function declarations are not order dependent
  - Similar to JS, no need to define a prototype before use like C/C++
- `go run *.go` is a better way of doing things instead of `go build ...; ./...`
  - The Go compiler actually finds the `main()` function in the `main` package
- Packages **NEED** to be stored in **different directories**
  - You can't have files specifying different `package` headings in the same directory
  - The `package` directory name, `package` header (in the code), and `import __packageName__` has to be equal to `__packageName__` or things will fail
- Exported functions are denoted by CapitalCase
- Variadic functions must be of single type
- `defer` can be used for post-function code injection