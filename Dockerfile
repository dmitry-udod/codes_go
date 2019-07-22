FROM golang:1.12-alpine as builder
RUN go version
RUN echo $GOPATH
RUN mkdir -p $GOPATH/src/github.com/dmitry-udod/codes_go
WORKDIR $GOPATH/src/github.com/dmitry-udod/codes_go
ADD go.mod .
ADD go.sum .
ADD . .
RUN ls -lah
RUN go build
RUN ls -lah