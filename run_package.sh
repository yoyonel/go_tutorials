#!/usr/bin/env bash

# argument:
# - $1: relativ path to root package directory (contain a 'src' directory)
# - $2: name of go project to run
# example: $ ./run_package.sh Packages shopping
# ┗ tree Packages         
# Packages
# └── src
#     └── shopping
#         ├── db
#         │   └── db.go
#         ├── main
#         │   └── main.go
#         └── pricecheck.go

source ./env.sh

DOCKER_VOLUME_PATH="$(realpath .)"

# command for building project
GO_CMD="go run $1/src/$2/main/main.go"
# GO_CMD=bash

# echo "GO_CMD: $GO_CMD"

docker run \
	-it \
	--rm \
	-v $DOCKER_VOLUME_PATH:/usr/src/myapp \
	-w /usr/src/myapp \
	-e "GOPATH=/usr/src/myapp/$1" \
	$GO_IMAGE \
	$GO_CMD