#! /usr/bin/env sh
# `go tool dist list` to list possible values for GOOS & GOARCH
# see "libc" array from `zig targets` output for possible values of zig -target
#ZIGCC="zig cc -target aarch64-linux-gnu.2.32"
#ZIGCC="zig cc -target aarch64-linux-musl"
ZIGCC="zig cc"
CC=$ZIGCC CXX=$ZIGCC CGO_ENABLED=1 go build "$@"
