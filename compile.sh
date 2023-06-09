#!/bin/bash

set -x

container=$(buildah from docker.io/library/golang@sha256:d78cd58c598fa1f0c92046f61fde32d739781e036e3dc7ccf8fdb50129243dd8)
echo "Container: $container"
buildah copy $container "account" ./account
buildah copy $container "list" ./list
buildah copy $container "login" ./login
buildah copy $container "Dockerfile" ./Dockerfile
buildah config --env GOPATH="" $container

echo " --- debug ---"
buildah run $container find .

buildah run $container /bin/sh -c 'cd account; go build .'
buildah run $container /bin/sh -c 'cd list; go build .'
buildah run $container /bin/sh -c 'cd login; go build .'

buildah config --workingdir . $container
mountpoint=$(buildah mount $container)
echo "Mountpoint: $mountpoint"
buildah run $container find .
find .
mkdir sewshul
cp $mountpoint/go/account/account ./sewshul/account
cp $mountpoint/go/list/list ./sewshul/list
cp $mountpoint/go/login/login ./sewshul/login
buildah unmount $mountpoint
buildah rm $container
