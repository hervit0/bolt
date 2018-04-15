# Bolt

## Overview

`Bolt` is a CLI tool, written in Go.

Why "bolt"? It's fast to type, you'll be like... a lightning bolt :smirk:.

## Prerequisites

For development only:
- [Docker](https://www.docker.com/)
- [dex](https://github.com/Driftrock/dex)

Install dependencies:
```
bash-4.4# dep ensure
```

Otherwise, nothing.

## Install it

Install it from outside
```
curl -s https://raw.githubusercontent.com/hervit0/bolt/master/bolt-install | sh
```

(or unreleased: `curl -s https://raw.githubusercontent.com/hervit0/bolt/develop/bolt-install.sh | sh`)

Install it from this repo (development mode)
```
$ ./install
```

## Use it

```
$ bolt
This is Bolt, try `bolt help` for more
```

## Docs

- [Docker Go starter](https://hub.docker.com/_/golang/)
- [Dockerfile best practices](https://docs.docker.com/v17.09/engine/userguide/eng-image/dockerfile_best-practices/#use-multi-stage-builds)
- [Golang dep](https://gist.github.com/subfuzion/12342599e26f5094e4e2d08e9d4ad50d)
- [Golang Command Line Runner](https://golang.org/pkg/os/exec/)
- [CLI Library Get started](https://github.com/urfave/cli#getting-started)

## Cheat Sheet

```
$ dep init
$ dep ensure -add github.com/pkg/errors
```
