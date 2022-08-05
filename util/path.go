package util

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func RootDir() string {
	_, file, _, _ := runtime.Caller(0)
	fmt.Println(file)
	fmt.Println(filepath.Dir(file))
	rootDir := filepath.Join(filepath.Dir(file), "../")
	fmt.Println(rootDir)
	return rootDir
}
