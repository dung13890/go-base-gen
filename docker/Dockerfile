FROM golang:1.22-alpine3.18 as base

# The latest alpine images don't have some tools like (`git` and `bash`).
RUN apk update && apk upgrade && \
    apk add --no-cache gcc libtool musl-dev make openssh

WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download

# install the air binary so we get live code-reloading when we save files
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.56.2
