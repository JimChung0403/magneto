FROM golang:1.10.4-alpine AS build-env
RUN apk add --no-cache --update git openssl bzr build-base \
    && go get -u github.com/kardianos/govendor && go get github.com/invite0403/magneto \
    && cd github.com/invite0403/magneto \
    && govendor sync && go build -o bin/magneto github.com/invite0403/magneto/cmd/magneto




MAINTAINER invite0403

