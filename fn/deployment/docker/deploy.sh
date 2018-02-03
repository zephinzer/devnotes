#!/bin/sh
CURR_DIR=$(dirname $0);
CALLED_WITH_COMMAND="$0 $@";
DEFAULT_NAMESPACE="fn";
FLAG_DOWN='--down';

_=$(command -v docker &>/dev/null);
if [ "$?" != "0" ]; then printf "Docker is required for this script to run. Install it and try again.\n"; exit 1; fi;
_=$(command -v docker-compose &>/dev/null);
if [ "$?" != "0" ]; then printf "DockerCompose is required for this script to run. Install it and try again.\n"; exit 1; fi;

if [ "${NAMESPACE}" = "" ]; then NAMESPACE="${DEFAULT_NAMESPACE}"; fi;
if [ -z "${CALLED_WITH_COMMAND##*${FLAG_DOWN}*}" ]; then ACTION_TYPE='shutdown';
else ACTION_TYPE='startup';
fi;
if [ "${ACTION_TYPE}" = "shutdown" ]; then
  docker stack rm "${NAMESPACE}";
else
  docker swarm init &>/dev/null;
  docker stack deploy --compose-file "${CURR_DIR}/docker-compose.yml" "${NAMESPACE}";
fi;