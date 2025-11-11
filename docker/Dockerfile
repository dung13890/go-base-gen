FROM golang:1.25-alpine as base

# The latest alpine images don't have some tools like (`git` and `bash`).
RUN apk update && apk upgrade && \
    apk add --no-cache gcc libtool musl-dev make openssh

WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v2.6
