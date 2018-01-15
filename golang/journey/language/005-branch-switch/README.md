# Branches in Golang (`switch`)

## Pre-requisites
See the [Setting up Go in README.md](../../README.md#setting-up-go) to get started.

## What's Covered
If and elses. See the source code at [`./branch-switch.go`](./branch-switch.go).

Covered are:

- switch

Build and run it with:

```bash
go run *.go
```

## Learnings

- No need for `break`s any longer!
  - No fallthrough unlike JS
- Weird `switch` with a semi-colon after before specifying the variable to be evaluated - it's called a short statement before a declaration (*`switch [STATEMENT];[DECLARATION] {...}`*)
- In contextless `switch`, you can specify an expression
- First passing expression is evaluated, think of it as endless `||` operators and only the first needs to evaluate to a truthy value