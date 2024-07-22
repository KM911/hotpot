package watcher

import (
	"fmt"
	"testing"
)

func TestFolder(t *testing.T) {
	for i, v := range Folders() {
		fmt.Println(i, v)
	}

}

func TestDir(t *testing.T) {
	// _, folder := fs.Dir(fs.HomeDirectory)
	// for i, v := range folder {
	// 	fmt.Println(i, v)
	// }
	foler := ListAllDir()
	for i := range foler {
		fmt.Println(foler[i])
	}
}
