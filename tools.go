// +build tools

package tools

import (
	_ "github.com/kevinburke/go-bindata/go-bindata"
	_ "golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow"
	_ "honnef.co/go/tools/cmd/staticcheck"
)
