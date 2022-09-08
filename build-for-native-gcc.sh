#! /usr/bin/env sh
# `go tool dist list` for list of possible values
CGO_ENABLED=1 go build "$@"
