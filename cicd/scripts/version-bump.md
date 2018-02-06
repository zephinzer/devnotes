# Version bump script for CI/CD pipelines

This script assumes execution from a directory where version tagging scripts are located at the relative path of `./scripts/versioning`. The scripts assumed to be there can be found at this GitHub repository: https://github.com/zephinzer/version-tagging-scripts

```bash
DEFAULT_GIT_USER="Version Bump Script";
DEFAULT_GIT_EMAIL="none@noemail.com";
TAG_MAJOR='[major]';
TAG_MINOR='[minor]';
TEMP_DIR_NAME='temp';
REPO_URL=https://user:${USER_ACCESS_TOKEN}@your.hostname.com/namespace/repository.git;
COMMIT_BRANCH="$(git rev-parse --abbrev-ref HEAD)";
COMMIT_MESSAGE="$(git log -1 --pretty=%B)"
COMMIT_SHA="$(git log -1 --pretty=%H)";

stat ./scripts/versioning/iterate &>/dev/null;
if [ "$?" = "0" ]; then
  printf "\e[1m\e[31mERR > iteration script for version bumping was not found.\e[0m\n";
  exit 1;
fi;

git clone -b ${COMMIT_BRANCH} "${REPO_URL}" "${TEMP_DIR_NAME}";
cd ./${TEMP_DIR_NAME};
git checkout ${COMMIT_BRANCH} .;

GIT_USERNAME="$(git config --global user.name)";
if [ "${GIT_USERNAME}" = "" ]; then
  GIT_USERNAME="${DEFAULT_GIT_USER}";
fi;
git config --global user.name "${GIT_USERNAME}";

GIT_EMAIL="$(git config --global user.email)";
if [ "${GIT_EMAIL}" = "" ]; then
  GIT_EMAIL="${DEFAULT_GIT_EMAIL}";
fi;
git config --global user.email "${DEFAULT_GIT_EMAIL}";

if [ -z "${COMMIT_MESSAGE##*${TAG_MAJOR}*}" ]; then
  printf "MAJOR VERSION BUMP\n";
  ./scripts/versioning/iterate major -q -i
elif [ -z "${COMMIT_MESSAGE##*${TAG_MINOR}*}" ]; then
  printf "MINOR VERSION BUMP\n";
  ./scripts/versioning/iterate minor -q -i
else
  printf "PATCH VERSION BUMP\n";
  ./scripts/versioning/iterate patch -q -i
fi;

git push origin "${COMMIT_BRANCH}" --tags;
cd ..;
rm -rf ./${TEMP_DIR_NAME};
```