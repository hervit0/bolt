# Bolt

## Overview

Bolt is a lightning fast CLI tool that speeds up development.

## Using Bolt CLI

At present, Bolt is just a CLI proxy which runs other CLI tools. For it to be useful you will need the following tools installed on your local workstation:
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- Git

### Installing Bolt

Visit the [releases page](https://github.com/hervit0/bolt/releases) for the latest version.

Run the following command to install bolt and start using it locally (remember to update the RELEASE_VERSION first):
```
RELEASE_VERSION="v0.1.7" sh -c 'curl -Lf https://github.com/hervit0/bolt/releases/download/$RELEASE_VERSION/bolt-darwin-amd64 --output bolt && chmod 755 bolt && mv bolt /usr/local/bin/bolt'
```

Then you should be able to try the following command to prove it is working

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
$ go test
```
