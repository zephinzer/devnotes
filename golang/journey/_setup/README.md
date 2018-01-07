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
Run the following to install all of Golang's toolchains:

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