#!/usr/bin/env bash

# url: https://hub.docker.com/_/golang/
GO_VERSION=1.6
GO_IMAGE=golang:$GO_VERSION

DOCKER_VOLUME_PATH="$(realpath .)"

# GoDoc http server
GODOC_HTTP_PORT=6060
DOCKER_GODOC_HTTP_PORT=6060
DOCKER_GODOC_CONTAINER=godoc_server
