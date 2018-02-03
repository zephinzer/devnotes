# Golang

So I started my journey in Golang on 7th January 2018. This is a collection of my notes and example applications to prove my understanding of Golang to myself.

The [`./language` directory](./language) contains code demonstrating language semantics.

The [`./application` directory](./application) contains implementation of concepts and other higher level things one can cover with Golang.

# Golang References
1. [Golang Language Specification](https://golang.org/ref/spec)
2. [Tour of Go](https://tour.golang.org/welcome/1) (**really useful**)

# Setting up Go

## Mac OS X

### Install Brew
Make sure you have homebrew:

```bash
brew -v
```

If you don't, install it via:

```bash
/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
```

Or if you're weary of strange commands in GitHub repos, the website for Homebrew can be found at [https://brew.sh/](https://brew.sh/).

### Install Go
Install Go by running:

```bash
brew install go
```

Verify the installation by running:

```bash
go version
```

You should see an output similar to:

```
go version go1.9.2 darwin/amd64
```

## Install Go extras

### Visual Studio Code
Run the following to install all of Golang's toolchains for VSCode integration:

```bash
go get -u -v github.com/nsf/gocode
go get -u -v github.com/rogpeppe/gopkgs
go get -u -v github.com/ramya-rao-a/go-outline
go get -u -v github.com/acroca/go-symbols
go get -u -v golang.org/x/tools/cmd/guru
go get -u -v golang.org/x/tools/cmd/gorename
go get -u -v sourcegraph.com/sqs/goreturns
go get -u -v github.com/golang/lint/golint
```

## Using Go

To build a Go application for export:

```bash
go build [...files.go]
```

To run a Go application in development:
```bash
go run [...files.go]
```
