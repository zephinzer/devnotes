# Exercuse #1: Hello World
This exercise covers:

1. Setting up a new function
2. Running a function in development
3. Verifying a function via contract testing
4. Updating a function by pushing to a Docker registry
5. Registering a new application
6. Deploying the function

## 1.1 Setting up a new function
Go to the `./src` directory, initialize your new function and navigate into it's directory:

```bash
# run this from the ./src directory
fn init --runtime node hello_world && cd hello_world
```

Open the generated `func.yaml` in an IDE. This is also known as the **function definition file**. Now let's add a `name` property set to `hello_world`, and a `path` property set to `/hello_world`. Your `func.yaml` should look like:

```yaml
name: hello_world
version: 0.0.1
runtime: node
entrypoint: node func.js
path: /hello_world
```

Note the `node func.js` in the `entrypoint` property. This means that this function will be executed by calling the Node runtime with the `func.js` script. Which we shall now create.

## 1.2 Running a function in development
Create the application code file:

```bash
touch func.js
```

Paste the following JS code inside:

```javascript
console.log('"hello world!"');
```

Test run application:

```bash
fn --verbose run
```

You should see a nice `"hello_world!"` output.

> Output is generated from printing to standard I/O. For Node, `console.log` and `console.info` does that. Alternatively, to go to a lower level, you could use `process.stdout.write(...)` too.

## 1.3 Verifying a function via contract testing
Let's write the tests for it. Fn comes with the ability to specify tests in a `test.json` file and have it run using an inbuilt command, `fn test`.

We proceed by creating a `test.json` file in the `hello_world` directory:

```bash
# run this in the ./src/hello_world directory
touch test.json
```

Paste the following JSON inside:

```json
{
  "tests": [
    {
      "input": {
        "body": ""
      },
      "output": {
        "body": "hello world!"
      }
    }
  ]
}
```

> This specifies we have one test where we send an input of an empty string and expect a response of `"hello world!"`, which as seen from above is truthy.

Run the tests using:

```bash
fn test
```

You should see that it passes. Congratuations (;

## 1.4 Updating a function by pushing to a Docker registry
Remember that we have a local throwaway Docker registry located at `localhost:5000`? We can push our application that Docker registry by running:

```bash
FN_REGISTRY=127.0.0.1:5000 fn build;
FN_REGISTRY=127.0.0.1:5000 fn push;
```

Check that it's properly pushed by querying the local Docker registry's API:

```bash
curl 127.0.0.1:5000/v2/_catalog
```

You should receive a JSON response with a `repositories` property containing an array with a `hello_world` item as follows:

```json
{"repositories":["hello_world"]}
```

To verify the version tag, run:

```bash
curl 127.0.0.1:5000/v2/hello_world/tags/list
```

This should yield a JSON response with a list of tags for the `hello_world` repository as follows:

```json
{"name":"hello_world","tags":["0.0.1"]}
```

Great, we're pushing it to the correct place!

## 1.5 Registering a new application
Every function necessarily belongs to an application. We proceed by creating the application. In this case it is `basic-application`. Create it with:

```bash
fn apps create basic-application
```

You should see the following response:

```
Successfully created app:  basic-application
```

Alternatively, if you've already deployed other functions to the `basic-application` application, you should see:

```
ERROR: App already exists
```

This is fine too, we just need the application to exist.

## 1.6 Deploying the function
Finally, run the following command to deploy our `hello_world` function to the `basic-application` application:

```bash
FN_REGISTRY=127.0.0.1:5000 fn deploy --app basic-application
```

Now check out your cool function at http://localhost:8080/r/basic-application/hello_world which should respond with... (*surprise surprise!*) `hello world!`.

## End of Exercise 1
Move on to [Exercise 2](./README-EX02.md).
