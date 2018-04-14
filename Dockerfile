FROM golang:1.10.1-alpine3.7

RUN apk update && apk upgrade && \
    apk add --no-cache bash
# Install dep for golang dependency management
# RUN curl -LO https://github.com/golang/dep/releases/download/v0.3.2/dep-linux-amd64 \
#   && mv dep-linux-amd64 /usr/bin/dep \
#   && chmod +x /usr/bin/dep

# COPY . /go/src/github.com/Driftrock/cdex
# WORKDIR /go/src/github.com/Driftrock/cdex

# RUN dep ensure
# RUN go-wrapper install
