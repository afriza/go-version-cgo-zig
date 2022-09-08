#! /usr/bin/env sh
# `go tool dist list` for list of possible values
CC=aarch64-linux-gnu-gcc CXX=aarch64-linux-gnu-cpp CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go build "$@"
