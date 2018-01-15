# Objects in Golang

## Pre-requisites
See the [Setting up Go in README.md](../../README.md#setting-up-go) to get started.

## What's Covered
If and elses. See the source code at [`./objects.go`](./objects.go).

Covered are:

- struct
- specifying self-referencing struct
- initialising a self-referencing struct

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
- Self-referencing structs can only do so using pointer concepts