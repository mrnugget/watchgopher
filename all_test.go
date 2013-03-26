package main

import (
	"path/filepath"
	"runtime"
)

var (
	_, filename, _, _ = runtime.Caller(0)
	fixtures          = filepath.Dir(filename) + "/fixtures"
)
