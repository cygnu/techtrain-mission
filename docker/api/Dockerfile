FROM golang:1.15.2

ENV GOPATH /go
ENV PATH $PATH:$GOPATH/bin

RUN mkdir -p /go/src/api
COPY . /go/src/api
WORKDIR /go/src/api
