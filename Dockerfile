FROM golang:1.12-alpine as builder
RUN go version
RUN echo $GOPATH
RUN mkdir -p $GOPATH/src/codes
WORKDIR $GOPATH/src/codes
ADD . .
RUN go build
RUN ls -lah