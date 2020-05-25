#!/bin/sh

set -e

SEM_VER=$(git describe --abbrev=0)
CODE_VER=$(git describe --dirty="-mod")
GIT_COMMIT=$(git rev-parse HEAD)

echo "Building container image"
echo "Version: " $CODE_VER
echo "Commit: " $GIT_COMMIT

docker build --force-rm --build-arg=VERSION=$CODE_VER --build-arg=VCS_REF=$GIT_COMMIT -f ./docker/Dockerfile -t akleinloog/http-logger:$CODE_VER .

#--build-arg=VERSION=$CODE_VER --build-arg=VCS-REF=$GIT_COMMIT

if [ $CODE_VER == $SEM_VER ]
  then
    echo "Publishing image as latest version to docker hub"
    docker tag akleinloog/http-logger:$CODE_VER akleinloog/http-logger:latest
    docker push akleinloog/http-logger:$CODE_VER
    docker push akleinloog/http-logger:latest
  else
    echo "This is not an official release, the image won't be pushed to docker hub"
fi

echo "Build completed"