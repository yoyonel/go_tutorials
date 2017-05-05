#!/usr/bin/env bash

echo "DOCKER_VOLUME_PATH: $DOCKER_VOLUME_PATH"
echo "GO_IMAGE: $GO_IMAGE"
echo "GO_CMD: $GO_CMD"

docker run \
	--rm \
	-v $DOCKER_VOLUME_PATH:/usr/src/myapp \
	-w /usr/src/myapp \
	$GO_IMAGE \
	$GO_CMD