# Ghost Deployment
This directory contains deployment manifests for the Ghost publishing platform.

Inside the deployment type directories, there is a `deploy.sh` file. You can use that to deploy the manifest in the recommended order.

- [Docker Deployment](./docker)

## Deployment Customisations
### Usage of MySQL
The deployment manifests assume usage of a MySQL database over a SQLite3 file so that we can keep the deployment stateless.

## Setting Up Ghost
After deployment is done, you'll need to setup your application. Locally, go to http://localhost:2368/ghost to do this.

## Relevant Links
- [Official Ghost Website](https://ghost.org/)
- [Ghost on GitHub](https://github.com/TryGhost/Ghost)
- [Ghost on DockerHub](https://hub.docker.com/_/ghost/)
