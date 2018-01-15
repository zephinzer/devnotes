# Arrays in Golang

## Pre-requisites
See the [Setting up Go in README.md](../../README.md#setting-up-go) to get started.

## What's Covered
If and elses. See the source code at [`./arrays.go`](./arrays.go).

Covered are:

- arrays
- nestedArrays

Build and run it with:

```bash
go run *.go
```

## Learnings

- Constructors use `{}` instead of `()`
- When using the `{}` constructor, arguments **must** appear in the order in which they were defined
- No more annoying `*` and `->` from C!
  - No more `*` for pointer value referencing
  - No more `->`, use the plain ol' `.`
- Nested arrays are ugly as hell and we need to use pointers. Is there a better way?
