#!/bin/bash

set -euo pipefail

# run the actual tests.
export GOVALUATE_TORTURE_TEST="true"
go test ./...
go test -bench=. -benchmem ./...

# coverage
# disabled by default because the benchmark run above is the historical script behavior.
#go tool cover -func=coverage.out
