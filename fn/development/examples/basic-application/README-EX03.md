# Exercise #3: Reflection
This exercise covers:

1. Runtime environment of a function
2. Retrieving the HTTP method
3. Retrieving other HTTP headers
4. Retrieving URL query parameters
5. Asynchronous functions
6. Logging to server but not to client

## Pre-requisites
Export the `FN_REGISTRY` variable as covered in previous exercises:

```bash
export FN_REGISTRY=127.0.0.1:5000
```

## Understanding the runtime environment
Create a new app called reflection:

```bash
# run this from the ./src directory
fn init --runtime node reflection && cd reflection
```

Add our application `name` and `path` as in previous examples (into `func.yaml`):

```yaml
name: reflection
...
path: /reflection
```

Now create the `func.js` with the following code:

```javascript
process.stdin.resume();
process.stdin.setEncoding('utf8');
console.log(process.env);
```

Now deploy it:

```bash
fn deploy --app basic-application;
```

Access your app via `curl`:

```bash
curl -X GET http://localhost:8080/r/basic-application/reflection
```

The reponse should be a JSON string that represents the environment of the function. This is where meta-data about how the function was invoked can be retrieved from.

## Retrieving the HTTP method
RESTful APIs depend on the HTTP verb, assuming we are building one (please don't unless you're either sure of what you're doing or it's a management decision) with a serverless architecture, we can derive the HTTP verb via the `FN_METHOD` environment variable.

## Retrieving other HTTP headers
Let's access our app and set some custom headers via `curl` again:

```bash
curl -H "WeIrd-Tok3n: ezpz1234" -X PATCH http://localhost:8080/r/basic-application/reflection
```

Inspect the response and you should see an environment variable `FN_HEADER_Weird-Tok3n` with a value of `"ezpz1234"`. Note that Fn assumes your HTTP headers are following the Capital-Dash case conventions, so a header with the key `hello-world` will become `FN_HEADER_Hello-World`.

## Retrieving URL query parameters
Let's access our app and set some custom query parameters via `curl` again:

```bash
curl -X GET "http://localhost:8080/r/basic-application/reflection?hello=world&hey=developers%20and%20engineers"
```

Unfortunately, Fn does not parse this for us. Let's draw our attention to the environment variable `FN_REQUEST_URL` which contains the full URL. We can see our URL query parameters there. Parsing it is up to us. For Node, we can use the in-built `url` library to do this:

```javascript
const url = require('url');
const paramsc = url.parse('http://localhost:8080/r/basic-application/reflection?hello=world&hey=developers%20and%20engineers').query.split('&').reduce((prev,curr) => {
    const param = curr.split('=', 2);
    prev[param[0]] = decodeURIComponent(param[1]);
    return prev;
  }, {});
```

## Asynchronous Functions
The main benefit of a serverless architecture is a non-blocking gradual degradation in your application. This is done through an internal queue and thus functions make the most sense when they are asynchronous.

Let's make our `reflection` function asynchronous. Go to the `func.yaml` and add a new line as follows:

```yaml
...
type: async
```

Let's deploy our application again and see the difference:

```bash
fn deploy --app basic-application
```

Now access it via `curl`:

```bash
curl -X GET http://localhost:8080/r/basic-application/reflection
```

Notice the instant response time and wait, where did our response go to? In such cases, how can we check on the function's details? Brings us to the next section:

## Logging to server but not to client
Go to the `func.js` file and append the following lines of code:

```javascript
// ...
console.error('it reached the end');
```

Deploy the function again:

```bash
fn deploy --app basic-application
```

Now `curl` the endpoint so that we are sure we got a log:

```bash
curl -X GET http://localhost:8080/r/basic-application/reflection
```

Take note of the returned `call_id`.

Now, access Docker services to get the ID of the logs server (in the developmnt deployment we are using the same database as the API):

```bash
docker stack services fn | grep fn_db | tr -s ' ' | cut -f 1 -d ' '
```

Access the database using:

```bash
mysql -uroot -pdefault_root_password -h127.0.0.1 -P3306 fn
```

Select all logs that contain the `call_id` from above. Assuming the `call_id` we received from the `docker stack services fn | ...` command is `12345678901234567890123456`, we run:

```mysql
SELECT * FROM logs WHERE log LIKE "%FN_CALL_ID: '12345678901234567890123456'%";
```

You should see the logs with a message: `"it reached the end"`!

Let's inspect the behaviour of the `console.error` now. Change the function back to synchronous by modifying the `func.yaml` so that the `type` property equals to `sync`:

```yaml
...
type: sync
```

Deploy it one last time:

```bash
fn deploy --app basic-application
```

Access it once more via `curl`:

```bash
curl -X GET http://localhost:8080/r/basic-application/reflection
```

Notice that the output from `console.error` is **NOT** shown. This is how we output logs that are only visible from the server-side and not the client side.

## End of Exercise 3
Move on to [Exercise 4](./README-EX04.md).