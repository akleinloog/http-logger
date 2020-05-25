# http-logger

A simple HTTP Logger in Go.

## Run a local instance

After cloning the repository, start a new local instance using
```
docker-compose up --build
```

## Build the image

The build script will automatically build and tag the image.
If it is a clean build, i.e., no local changes, and no commits since the last version tag, the image is automatically pushed to docker hub.

Build a new image using
```
./deploy.sh
```

