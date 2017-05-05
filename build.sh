#!/usr/bin/env bash

# argument:
# - $1: path to (Go) src directory

source ./env.sh

# command for building project
GO_CMD="go build -v"

DOCKER_VOLUME_PATH="$(realpath $1)"

docker run \
	--rm \
	-v $DOCKER_VOLUME_PATH:/usr/src/myapp \
	-w /usr/src/myapp \
	$GO_IMAGE \
	$GO_CMD