# Basic Fn Application Structure
This is an example project for an Fn serverless application.

Before running anything in here, make sure you've a local Fn deployment. Instructions can be found in the [`deployment` directory](../../../deployment).

To check if everything is up correctly, run:

```bash
docker stack services fn
```

All rows should have `1/1` in the `REPLICAS` column.

You should also have `fn` installed.

For OS X:

```bash
brew update && brew install fn;
```

For other Linux-based systems:

```bash
curl -LSs https://raw.githubusercontent.com/fnproject/cli/master/install | sh
```

## Directory Structure

### `./scripts`
This directory contains scripts which can be used in a Continuous Integration/Delivery (CI/CD) pipeline to deploy code.

### `./src`
Each directory in `./src` is one independently deployable function.

### `./test`
This directory contains examples of the various types of application-level verification we can do for applications built on a serverless architecture. Function-level verification is done in the function directory itself, and will be covered below.

- - -

## Exercises

- [1. Hello World](./README-EX01.md)
  - Setting up a new function
  - Running a function locally
  - Testing a function using Fn's in-built contract tests
  - Updating a function by pushing to a Docker registry
  - Deploying a function to a Fn server
- [2. Addition](./README-EX02.md)
  - Development environment provisioning
  - Test-driven-development in functions
  - Reading from input
- [3. Reflection](./README-EX03.md)
  - Digging into the runtime environment
  - Error logging
