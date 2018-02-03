#!/bin/sh
CURR_DIR=$(dirname $0);
CALLED_WITH_COMMAND="$0 $@";

ACTION="start";
FLAG_DOWN='--down';
FLAG_STOP='--down';

STACK_NAME='ironfx-dev';

_=$(command -v docker);
if [ "$?" != "0" ]; then
  printf "\e[31m\e[1mThis script requires Docker to be installed.\e[0m\n";
  exit 1;
fi;

docker swarm init &>/dev/null;

if [ -z "${CALLED_WITH_COMMAND##*${FLAG_DOWN}*}" ]; then
  ACTION="stop";
elif [ -z "${CALLED_WITH_COMMAND##*${FLAG_STOP}*}" ]; then
  ACTION="stop";
fi;

if [ "${ACTION}" = "stop"]; then
  docker stack rm "${STACK_NAME}";
else
  docker stack deploy --compose-file "${CURR_DIR}/docker-compose.yml" "${STACK_NAME}";
  printf "> run \`docker stack ps ironfx-dev\` to see the status of the stack.\n";
  printf "> run \`docker service ls\` to see the status of the services.\n";
fi;
