#!/usr/bin/env bash

set -euo pipefail

die() {
    echo "$@" >&2
    exit 1
}

case "$PROTOBUF_VERSION" in
2*)
    basename=protobuf-$PROTOBUF_VERSION
    ;;
3*)
    basename=protobuf-cpp-$PROTOBUF_VERSION
    ;;
*)
    die "unknown protobuf version: $PROTOBUF_VERSION"
    ;;
esac

curl -sL https://github.com/google/protobuf/releases/download/v$PROTOBUF_VERSION/$basename.tar.gz | tar zx

cd protobuf-$PROTOBUF_VERSION

./configure --prefix=/home/travis && make -j2 && make install
