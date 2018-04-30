#!/bin/sh
# ref: https://developer.github.com/v3/repos/releases/#get-the-latest-release
USER="$1";
REPO="$2";
REPO_PATH="${USER}/${REPO}";
printf -- "Retrieving latest tag for ${REPO_PATH}... ";
RELEASE="$(curl -sSL -X GET "https://api.github.com/repos/${REPO_PATH}/releases/latest")";
RELEASE_TAG="$(printf -- "${RELEASE}" | jq -r '.tag_name')";
if [ "${RELEASE_TAG}" = "null" ]; then
  printf -- "${RELEASE}" | jq -r '.message';
fi;
