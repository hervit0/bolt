# Bolt

## Overview

`Bolt` is a CLI tool, written in Go.

## Prerequisites

[Docker](https://www.docker.com/) and [dex](https://github.com/Driftrock/dex)
```
dep ensure
```

## Run it

Inside Docker
```
go build bolt.go
```

Outside Docker
```
env GOOS=darwin GOARCH=amd64 go build -v bolt.go
```

Use it
```
$ ./bolt
boom boom! No way it's working!
```

## Docs

- [Docker Go starter](https://hub.docker.com/_/golang/)
- [Dockerfile best practices](https://docs.docker.com/v17.09/engine/userguide/eng-image/dockerfile_best-practices/#use-multi-stage-builds)
- [Golang dep](https://gist.github.com/subfuzion/12342599e26f5094e4e2d08e9d4ad50d)
- [Golang Command Line Runner](https://golang.org/pkg/os/exec/)

## Cheat Sheet

```
$ dep init
$ dep ensure -add github.com/pkg/errors
```
