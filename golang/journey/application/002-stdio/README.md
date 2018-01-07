# FMT Package in Golang

## Pre-requisites
See the [Setting up Go README.md](../../_setup/README.md) to get started.

## What's Covered
A basic hello world in Golang! See the source code at [`./stdio.go`](./stdio.go).

Build it with:

```bash
go run stdio.go
```

## Learnings

- Errorf prints to the error interface, to see it, run:
  ```bash
  go run stdio.go 2&> tee
  ```