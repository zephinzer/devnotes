# Exercise #2: Addition
This exercise covers:

1. Provisioning the development environment
2. Test-driven-development in functions
3. Reading user input

## 2.1 Provisioning the development environment
To make our lives easier, we avoid specifying `FN_REGISTRY` at every command. We can export it as a global environment variable using:

```bash
export FN_REGISTRY=127.0.0.1:5000
```

> What about deploying to production? Well, **never deploy using your local machine**, leave that to the CI/CD pipeline!

## 2.2 Test-driven-development in functions
To demonstrate writing of tests first for a function, we proceed by creating a test contract first. Run the following to initialise a new function.

```bash
# run this from the ./src directory
fn init --runtime node addition && cd addition
```

Create the `test.json`:

```bash
touch test.json
```

Paste the following JSON into the `test.json`:

```json
{
  "tests": [
    {
      "input": {
        "body": [1, 2, 3]
      },
      "output": {
        "body": 6
      }
    },
    {
      "input": {
        "body": [-1, 2, 3]
      },
      "output": {
        "body": 4
      }
    }
  ]
}
```

> This `test.json` specifies two tests, the first test indicates that an array of `1, 2, 3` is the input and the result should be a summation of those numbers, `6`. The second test indicates the same except for the `1` which is now negated, resulting in a final result of `4`.

Run the tests using:

```bash
fn test
```

We should get an error that looks like:

```
...
Error: Cannot find module '/function/func.js'
...
```

This is perfectly fine. Now let's get the tests to pass.

## 2.3 Reading user input
Let's create the code for it by creating a file named `func.js` and placing the following code into it:

```javascript
process.stdin.resume();
process.stdin.setEncoding('utf8');
process.stdin.on('data', (_data) => {
  const data = JSON.parse(_data);
  console.log(
    data.reduce((prev, curr) => (prev + curr), 0)
  );
});
```

As you can see, we're reading from `process.stdin` which is how functions receive the data intended for it. Then we proceed by reducing the array of numbers from the above `test.json` and returning the result through `console.log`, which is what the user will receive as a response.

Now, run the tests again, and we should be getting two greens (two test cases from `test.json`):

```bash
fn test
```

## Deploying the code
Once again, add a `name` and `path` property to the `func.yaml` function definition file:

```yaml
name: addition
...
path: /addition
```

Then run:

```bash
fn deploy --app basic-application
```

Now access it by `curl`-ing the URL:

```bash
curl -d "[2,2,3]" -X POST http://localhost:8080/r/basic-application/addition
```

You should see `7` as the response!

## End of Exercise 2
Move on to [Exercise 3](./README-EX03.md).