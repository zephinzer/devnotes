#!/bin/sh
docker swarm init &>/dev/null;

docker stack deploy \
  --compose-file ./docker-compose.yml \
  ghost;