#!/bin/sh

set -e

echo "Building Container"

SEMVER=$(git describe --abbrev=0)
CODEVER=$(git describe --dirty="*")
GIT_COMMIT=$(git rev-parse HEAD)

echo "Version: " $CODEVER
echo "Commit: " $GIT_COMMIT

if [ $CODEVER == $SEMVER ]
  then
    echo "Publishing Image"
  else
    echo "Not an official release"
fi