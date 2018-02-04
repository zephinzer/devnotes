# Exercise #5: Testing in Production

This exercise covers:

1. Prometheus metrics collection
2. Number of times each function was called
3. Errored functions
4. Processing duration

## Prometheus metrics collection
Navigate to [http://localhost:9090](http://localhost:9090) which is where Prometheus resides assuming you've [followed the instructions from the `deployment` directory](../../../deployment/docker).

Inside Prometheus, navigate to the Graph tab from the header navbar. Enter the following into the *Expression* textbox:

```promql
fn_calls{fn_appname="basic-application"}
```

This queries all calls to the `basic-application` application. If you've just finished the exercises, the graph should be populated. You could also run the following block of commands to simulate some calls using the completed functions in this repository. But before that, if you haven't created an `app.yaml`, rename the `app_completed.yaml` to `app.yaml` and run `fn deploy --all` from the `./src` directory. When that is done, you may run the following:

```bash
# run this from your terminal, not Prometheus
curl -X POST http://localhost:8080/r/basic-application/hello_world_completed & curl -d "[11,22,33]" -X POST http://localhost:8080/r/basic-application/addition_completed & curl -d "[59,273,88]" -X POST http://localhost:8080/r/basic-application/addition_completed & curl -X POST http://localhost:8080/r/basic-application/reflection_completed & curl -X PUT http://localhost:8080/r/basic-application/reflection_completed & curl -X GET http://localhost:8080/r/basic-application/reflection_completed;
```

Note that we are calling `hello_world_completed` 1 time, `addition_completed` 2 times, and `reflection_completed` 3 times.


## Number of times each function was called

Enter the following into the *Expression* textbox again and hit **Execute**:

```promql
fn_calls{fn_appname="basic-application"}
```

You should see the number of calls made to each function.

> This metric is useful to see the potential impact updating a function may have. If many other functions depend on this one function, it might be better to play it safe and deploy a breaking version elsewhere. Alternatively, for function deprecation, if the function is no longer called and hasn't been called for awhile, it would be safe to delete the function.

## Errored functions

Let us create an error now by intentionally providing unexpected data to the `addition` function. This will cause the function to error out because the data cannot be split:

```bash
curl -vv -d 'a' -X POST http://localhost:8080/r/basic-application/addition_completed;
```

Wait till the `curl` process completes. You should receive a `HTTP/1.1 502 Bad Gateway` with an error response body. Now navigate to the Prometheus graph once more at http://localhost:9090/graph.

Enter the following PromQL:

```promql
fn_errors{fn_appname="basic-application"}
```

You should see that the error count has incremented.

> This metric is useful for testing in production when we deploy a new version of a function. If the function has always run without errors but is suddenly doing so, we can automate the revert via a Grafana/Prometheus alert which triggers a webhook that rolls back to a previous working version. Neat.

## Processing duration

Enter the following PromQL:

```promql
fn_container_busy_duration_seconds_sum{fn_appname="basic-application"} / fn_container_busy_duration_seconds_count{fn_appname="basic-application"}
```

This shows us the processing duration of each function so that we can optimise resources to fit our service level agreements.

