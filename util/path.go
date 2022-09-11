package util

import (
	"path/filepath"
	"runtime"
)

// RootDir returns the root directory(repository root) path.
// However, when executing binary returns the 1 level higher path for binary .
func RootDir() string {
	_, file, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(file), "../")
}
