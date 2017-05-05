#!/usr/bin/env bash

source ./env.sh

# go command to launch local http server for doc (at 6060 port)
GO_CMD="godoc -http=:$GODOC_HTTP_PORT"

docker run \
	--rm -d \
	--name $DOCKER_GODOC_CONTAINER \
	-p $GODOC_HTTP_PORT:$DOCKER_GODOC_HTTP_PORT \
	$GO_IMAGE \
	$GO_CMD