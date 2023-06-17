package internal

import (
	"path/filepath"
	"runtime"
)

var _, b, _, _ = runtime.Caller(0)

// Root is the root directory of the project.
var Root = filepath.Join(filepath.Dir(b), "../")
