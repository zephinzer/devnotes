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
  - Running a function in development
  - Verifying a function via contract testing
  - Updating a function by pushing to a Docker registry
  - Registering a new application
  - Deploying the function
- [2. Addition](./README-EX02.md)
  - Provisioning the development environment
  - Test-driven-development in functions
  - Reading user input
- [3. Reflection](./README-EX03.md)
  - Runtime environment of a function
  - Retrieving the HTTP method
  - Retrieving other HTTP headers
  - Retrieving URL query parameters
  - Asynchronous functions
  - Logging to server but not to client
- [4. Continuous Integration](./README-EX04.md)
  - Registering an application by directory
  - Testing all functions in a pipeline
  - Deploying all functions from the pipeline
  - Moving from integration to delivery
- [5. Testing in Production (WIP)](./README-EX05.md)
  - Prometheus metrics collection
  - Number of times each function was called
  - Errored functions
  - Processing duration