#!/bin/bash

set -x

container=$(buildah from docker.io/library/golang@sha256:d78cd58c598fa1f0c92046f61fde32d739781e036e3dc7ccf8fdb50129243dd8)
echo "Container: $container"
buildah copy $container "*" .
buildah config --env GOPATH="" $container
buildah run $container cd account
buildah run $container go mod download
buildah run $container go build .
buildah run $container cd ../list
buildah run $container go mod download
buildah run $container go build .
buildah run $container cd ../login
buildah run $container go mod download
buildah run $container go build .
mountpoint=$(buildah mount $container)
echo "Mountpoint: $mountpoint"
buildah run $container pwd
buildah run $container find .
find .
mkdir sewshul
cp $mountpoint/go/sewshul ./sewshul/sewshul
buildah unmount $mountpoint
buildah rm $container