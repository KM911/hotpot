package watcher

import (
	"fmt"
	"testing"

	"github.com/KM911/fish/fs"
)

func TestFolder(t *testing.T) {
	for i, v := range Folders() {
		fmt.Println(i, v)
	}

}

func TestDir(t *testing.T) {
	_, folder := fs.Dir(fs.HomeDirectory)
	for i, v := range folder {
		fmt.Println(i, v)
	}
}
