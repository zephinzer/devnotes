# Exercise #4. Continuous Integration (CI)

This exercise covers:

1. Registering an application by directory
2. Testing all functions in a pipeline
3. Deploying all functions from the pipeline
4. Moving from integration to delivery

## Pre-requisites
Export the `FN_REGISTRY` variable as covered in previous exercises:

```bash
export FN_REGISTRY=127.0.0.1:5000
```

## Registering an application by directory
Navigate to the `./src` directory and create an `app.yaml` file with the following contents:

```yaml
name: basic-application
```

This indicates that the `./src` directory is now the root for the application `basic-application`, enabling us to use the `--all` flag with `fn` commands.

## Testing all functions in a pipeline
CI pipelines run automated tests on the code to make sure that the code works. With Fn, the command is `fn test` to test an individual function.

In `./src`, run the following to test all functions we have created in previous exercises:

```bash
fn test --all
```

## Deploying all functions from the pipeline
Similar to the `fn test` command, the `fn deploy` command also has an `--all` flag capability. Run it with:

```bash
fn deploy --all
```

## Moving from integration to delivery
There are two ways of doing this, both with their merits and drawbacks. The first is through having two Fn servers, one to integrate, and the other to deliver. The second is through injecting a different app.yaml with the application name set to the commit hash which we can then promote to production.

The benefit of the first is that the system is completely separate and any major faults does not affect the production system. The cons is having to keep the staging system as similar to the production one. However, given that we can (and should) define our infrastructure as code, this risk is largely mitigated if one has the financial capabilities to do so. We can achieve this by setting the `API_URL` environment variable to a different deployment of Fn.

The benefit of the second is that we can test it directly in the production environment. With a serverless architecture using Fn, updating of the production application becomes a matter of changing the version number of the image to the same one as the commit hash. We can do this by running the deployer in a container that overwrites the value in `app.yaml`. An example script:

```bash
printf "-staging" >> ./src/app.yaml;
# continue as per-normal
# ...
```

The third way is to deploy directly to production which isn't as scary as it was before since we are deploying functions and not an application. The ability to deploy in parts mitigates the risk of an entire application breaking and allows us to easily perform a different type of testing that works based on proper monitoring systems and alerts, and the ability to quickly roll back to a previous version if needed. In the case of Fn, reverting a deployment is as simple as updating the version tag of the function image.

## End of Exercise 4
Move on to [Exercise 5](./README-EX05.md).