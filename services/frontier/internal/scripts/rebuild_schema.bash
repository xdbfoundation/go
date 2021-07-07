#! /usr/bin/env bash
set -e

# This scripts rebuilds the latest.sql file included in the schema package.
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
GOTOP="$( cd "$DIR/../../../../../../../.." && pwd )"

go generate github.com/xdbfoundation/go/services/frontier/internal/db2/schema
go generate github.com/xdbfoundation/go/services/frontier/internal/test
go install github.com/xdbfoundation/go/services/frontier
