# Fn Deployment on Docker Swarm

This directory contains a working `docker-compose` which you can use to spin up a production grade Fn locally. Simply run `./deploy.sh` to spin up a local single-node cluster. But before that...

## Pre-Deploy Setup

Due to `git` being unable to accomodate an empty directory, you'll need to create a directory in the `data` directory because MySQL will complain if the directory is non-empty:

- `./data/mysql/mnt`

## Usage

### What should be up
After running `./deploy.sh`, you should now have the following running:

- MySQL server on port 3306
- Redis on port 6379
- Fn server on port 8080
- Fn UI on port 4000
- Docker Registry on port 5000
- Prometheus on port 9090
- Grafana dashboard on port 3000

### Making sure it's all up
The stack may take awhile to completely spin up. If you're the impatient kind and need to see some progress, run:

```bash
docker stack ps fn
```

This will show you the deployment status of all services at the cluster level. To see it individually:

```bash
docker stack ps fn | grep %SERVICE_NAME%
```

To see the status of each service at the application level, run:

```bash
docker stack services fn
```



## Teardown
To shut down the cluster, run `./deploy.sh --down`. You may need to manually remove the network `fn_default` if it refuses to delete the network.

## Going to Production

### Separate the LogStore and API DB
In development, we use the same database for the API and the logs. In production, this should be different.

## Debugging

### `fn_default` network not found
- run `docker network ls` and search for `fn_default` in the `NAME` column
- copy the NETWORK ID of that row and run `docker network rm <NETWORK_ID>`