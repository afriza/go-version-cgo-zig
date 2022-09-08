#! /usr/bin/env sh
set -x

./build-for-linux-arm64-gcc.sh -o out-gcc "$@"
./build-for-linux-arm64-zig.sh -o out-zig "$@"

./build-for-native-gcc.sh -o out-n-gcc "$@"
./build-for-native-zig.sh -o out-n-zig "$@"

go version -m out-*
