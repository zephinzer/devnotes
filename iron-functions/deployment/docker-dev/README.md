# Docker Swarm Deployment for Iron Functions
The `./docker-compose.yml` specifies a stack deployment for deploying an Iron Functions meant for local development.

The following will be made available:

- Iron Functions API at `localhost:8080`
- Iron Functions UI at `localhost:4000`
- Docker Registry at `localhost:5000`

The local Docker Registry is useful for hastening the speed of deploying to a registry so that the image can be pulld by Iron Functions.

- - -

To start the stack, run:

```bash
./.deploy.sh
```

To stop the stack, run:

```bash
./.deploy.sh --down
```