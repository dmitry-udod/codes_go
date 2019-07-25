FROM golang:1.12-alpine as builder
RUN go version
ENV GO111MODULE=on
RUN echo $GOPATH
RUN mkdir -p $GOPATH/src/github.com/dmitry-udod/codes_go
WORKDIR $GOPATH/src/github.com/dmitry-udod/codes_go
ADD . .

# Run checks
RUN mkdir -p /app
ADD ./CHECKS /app

# Install additional packages
RUN apk add git
RUN apk add bash
RUN apk add curl

# build app
RUN go build