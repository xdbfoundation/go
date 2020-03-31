#! /bin/bash
set -e

printf "Running go vet...\n"
go vet -all -composites=false -unreachable=false -tests=false ./...

printf "Running go vet shadow...\n"
# Running shadown through "go run" is way too slow, so we need to install it locally
go get golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow

go vet -vettool=$(which shadow) ./...
