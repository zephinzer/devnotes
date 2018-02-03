#!/bin/sh
CURR_DIR=$(dirname $0);
CALLED_WITH_COMMAND="$0 $@";

FLAG_DOCKERFILE='--file';
FLAG_DOCKER_COMPOSE='--compose';
FLAG_DOCKER_SWARM='--stack';
FLAG_DOWN='--down';

STACK_NAME="mysql";

if [ -z "${CALLED_WITH_COMMAND##*${FLAG_DOCKERFILE}*}" ]; then
  docker build -t mysql:eg .
  docker run -v "$(pwd)/data/mysql:/var/lib/mysql" -p "3306:3306" mysql:eg
elif [ -z "${CALLED_WITH_COMMAND##*${FLAG_DOCKER_COMPOSE}*}" ]; then
  docker-compose \
    --file "${CURR_DIR}/docker-compose.yml" \
    up;
elif [ -z "${CALLED_WITH_COMMAND##*${FLAG_DOWN}*}" ]; then
  docker stack rm "${STACK_NAME}";
elif [ -z "${CALLED_WITH_COMMAND##*${FLAG_DOCKER_SWARM}*}" ]; then
  docker swarm init &>/dev/null;
  docker stack deploy \
    --compose-file "${CURR_DIR}/docker-compose.yml" \
    "${STACK_NAME}";
else
  printf "specify a flag to use:\n";
  printf "  --file    : run the Dockerfile\n";
  printf "  --compose : deploy the docker-compose.yml\n";
  printf "  --stack   : deploy to a local Docker Swarm\n";
  printf "  --down    : remove the application from the local Docker Swarm\n";
fi;