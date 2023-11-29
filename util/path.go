package util

import (
	"os"
	"path/filepath"
)

func CmdPath() string {
	dir, _ := os.Getwd()
	return dir
}

func ExecutePath() string {
	path, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(path)
}
