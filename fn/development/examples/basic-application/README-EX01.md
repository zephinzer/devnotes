# Exercuse #1: Hello World
This exercise covers:

1. Setting up a new function
2. Running a function locally
3. Testing a function using Fn's in-built contract tests
4. Updating a function by pushing to a Docker registry
5. Deploying a function to a Fn server

## Setting Up
Go to the `./src` directory, initialize your new function and navigate into it's directory:

```bash
# run this from the ./src directory
fn init --runtime node hello_world && cd hello_world
```

Open the generated `func.yaml` and add a `name` property set to `hello_world` and a `path` property set to `/hello_world`:

```yaml
name: hello_world
...
path: /hello_world
```

## Development
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

You should see a nice `hello_world!` output.

## Verification
Let's write the tests for it. Create a `test.json` file in the `hello_world` directory:

```bash
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

Run the tests using:

```bash
fn test
```

## Deployment

### Update the container registry
Remember that we have a local throwaway Docker registry located at `localhost:5000`. We can push our application that Docker registry by running:

```bash
FN_REGISTRY=127.0.0.1:5000 fn build;
FN_REGISTRY=127.0.0.1:5000 fn push;
```

Check that it's properly pushed by querying the local Docker registry's API:

```bash
curl 127.0.0.1:5000/v2/_catalog
```

You should receive a JSON response with a `repositories` property containing an array with a `hello_world` item.

```json
{"repositories":["hello_world"]}
```

To verify the version tag, run:

```bash
curl 127.0.0.1:5000/v2/hello_world/tags/list
```

You should receive a JSON response with a list of tags for the `hello_world` repository:

```json
{"name":"hello_world","tags":["0.0.1"]}
```

### Register the application
Now that our registry contains a version of our code, let's create the application:

```bash
fn apps create basic-application
```

You should see:

```
Successfully created app:  basic-application
```

Alternatively, if you've already deployed other functions to the `basic-application` application, you should see:

```
ERROR: App already exists
```

This is fine too, we just need the application to exist.

### Deploy the function
Finally, run the following command to deploy our `hello_world` function to the `basic-application` application:

```bash
FN_REGISTRY=127.0.0.1:5000 fn deploy --app basic-application
```

Now check out your cool function at http://localhost:8080/r/basic-application/hello_world