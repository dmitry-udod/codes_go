FROM golang:1.12-alpine as builder
RUN go version
RUN echo $GOPATH
RUN mkdir -p $GOPATH/src/github.com/dmitry-udod/codes_go
WORKDIR $GOPATH/src/github.com/dmitry-udod/codes_go
ADD . .
RUN go build
RUN ls -lah