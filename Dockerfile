FROM golang:1.12-alpine as builder
RUN go version
ENV GO111MODULE=on
RUN echo $GOPATH
RUN mkdir -p $GOPATH/src/github.com/dmitry-udod/codes_go
WORKDIR $GOPATH/src/github.com/dmitry-udod/codes_go
ADD . .
RUN ls -lah
RUN apk add --no-cache git
RUN go build
RUN ls -lah