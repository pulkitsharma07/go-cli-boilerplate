FROM golang:1.12-alpine3.10 AS builder

# git is required for fetching the dependencies.
# make is required for triggering the Makefile
# build-base is required for install gcc, which is required for running tests (go test)
RUN apk update && apk add --no-cache git && apk add --no-cache make && apk add --no-cache build-base

WORKDIR cli/

# Copying dependencies first to ensure that
# dependencies are not installed on each run.
COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go mod verify

COPY . .

RUN ./build_dist.sh

