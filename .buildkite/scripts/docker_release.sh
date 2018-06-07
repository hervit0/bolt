#!/bin/bash
set -x #echo on
set -e #exit on error

DARWIN_RELEASE_NAME=bolt-darwin-amd64

GOARCH=amd64 GOOS=darwin go build -v -o .build/${DARWIN_RELEASE_NAME}

go get github.com/aktau/github-release

github-release release \
    --user hervit0 \
    --repo bolt \
    --tag ${RELEASE_VERSION}

github-release upload \
    --user hervit0 \
    --repo bolt \
    --tag ${RELEASE_VERSION} \
    --name "${DARWIN_RELEASE_NAME}" \
    --file .build/${DARWIN_RELEASE_NAME}


# build all
# for GOOS in darwin linux; do
#   for GOARCH in 386 amd amd64; do
#     GOARCH=$GOARCH GOOS=$GOOS go build -v -o build/cdex-$GOOS-$GOARCH
#   done
# done