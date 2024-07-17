#!/bin/env bash

cd $(dirname $0)

bun build index.ts --compile --target=bun-linux-x64-modern --outfile dist/hinta-tarkkailija-linux-x64
bun build index.ts --compile --target=bun-linux-arm64 --outfile dist/hinta-tarkkailija-linux-arm64
