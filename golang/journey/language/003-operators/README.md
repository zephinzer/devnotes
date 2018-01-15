# Operators in Golang

## Pre-requisites
See the [Setting up Go in README.md](../../README.md#setting-up-go) to get started.

## What's Covered
See the source code at [`./branch-if-else.go`](./branch-if-else.go).

Covered are:

- <
- >
- <=
- >=
- !=
- ==
- ++
- --
- +
- +=
- -
- -=
- &
- &=
- |
- |=
- >>
- <<

Build and run it with:

```bash
go run *.go
```

## Learnings

- `d++` is an invalid expression, it's a statement unlike in other C-like languages
  - And so is `d+=1` and other state modifying stuff
- There is no pre-increment in Golang (ie `++d` is invalid)
