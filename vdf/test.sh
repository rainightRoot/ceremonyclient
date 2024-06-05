#!/bin/bash
set -euxo pipefail

SCRIPT_DIR="${SCRIPT_DIR:-$( cd "../$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )}"
ROOT_DIR="$SCRIPT_DIR"

BINDINGS_DIR="$ROOT_DIR/vdf"
BINARIES_DIR="$ROOT_DIR/target/release"

go generate

# Test the generated bindings
pushd "$BINDINGS_DIR" > /dev/null
LD_LIBRARY_PATH="${LD_LIBRARY_PATH:-}:$BINARIES_DIR" \
	CGO_LDFLAGS="-lvdf -L$BINARIES_DIR -ldl" \
	CGO_ENABLED=1 \
	LC_RPATH="$BINARIES_DIR" \
  go test
