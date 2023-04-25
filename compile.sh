#!/bin/bash

set -x

container=$(buildah from docker.io/library/golang@sha256:d78cd58c598fa1f0c92046f61fde32d739781e036e3dc7ccf8fdb50129243dd8)
echo "Container: $container"
buildah copy $container "*" .
buildah config --env GOPATH="" $container

#echo " --- Build account ---"
#buildah config --workingdir ./account $container
#buildah run $container go mod download
#buildah run $container go build .

#echo " --- Build list ---"
#buildah config --workingdir ./list $container
#buildah run $container go mod download
#buildah run $container go build .

#echo " --- Build login ---"
#buildah config --workingdir ./login $container
#buildah run $container go mod download
#buildah run $container go build .

#buildah config --workingdir . $container
#mountpoint=$(buildah mount $container)
#echo "Mountpoint: $mountpoint"
#buildah run $container pwd
buildah run $container find .
#find .
#mkdir sewshul
#cp $mountpoint/go/account ./sewshul/account
#cp $mountpoint/go/list ./sewshul/list
#cp $mountpoint/go/login ./sewshul/login
#buildah unmount $mountpoint
#buildah rm $container
