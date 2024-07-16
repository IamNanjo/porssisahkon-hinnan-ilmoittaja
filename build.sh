#!/bin/env bash

cd $(dirname $0)

bun build index.ts --compile --outfile dist/hinta-tarkkailija
