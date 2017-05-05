#!/usr/bin/env bash

# argument:
# - $1: path to (Go) src file

source ./env.sh

DOCKER_VOLUME_PATH="$(realpath .)"

# command for building project
GO_CMD="go run $@"

echo "GO_CMD: $GO_CMD"

docker run \
	--rm \
	-v $DOCKER_VOLUME_PATH:/usr/src/myapp \
	-w /usr/src/myapp \
	$GO_IMAGE \
	$GO_CMD