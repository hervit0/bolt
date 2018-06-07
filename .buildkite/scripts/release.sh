#!/bin/bash
set -x #echo on
set -e #exit on error

export RELEASE_VERSION=$(buildkite-agent meta-data get release-version)

docker-compose -f docker-compose.ci.yml build
docker-compose -f docker-compose.ci.yml run release