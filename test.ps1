$ErrorActionPreference = "Stop"

$env:GOVALUATE_TORTURE_TEST = "true"
go test ./...
go test -bench=. -benchmem ./...

# Coverage is disabled by default because the benchmark run above is the
# historical script behavior.
# go test -coverprofile coverage.out ./...
# go tool cover -func=coverage.out
