#!/bin/bash
set -x #echo on
set -e #exit on error

DARWIN_RELEASE_NAME=bolt-darwin-amd64

GOARCH=amd64 GOOS=darwin go build -v -o .build/${DARWIN_RELEASE_NAME}

TAG=v0.1.0

git remote -v

git tag ${TAG} 
git push --tags

go get github.com/aktau/github-release

github-release upload \
    --user danbadge \
    --repo bolt \
    --tag ${TAG} \
    --name "${DARWIN_RELEASE_NAME}" \
    --file build/${DARWIN_RELEASE_NAME}


# build all
# for GOOS in darwin linux; do
#   for GOARCH in 386 amd amd64; do
#     GOARCH=$GOARCH GOOS=$GOOS go build -v -o build/cdex-$GOOS-$GOARCH
#   done
# done