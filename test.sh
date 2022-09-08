#! /usr/bin/env sh
set -x

./build-for-linux-arm64-gcc.sh -o out-gcc "$@"
./build-for-linux-arm64-zig.sh -o out-zig "$@"

go version -m out-*
