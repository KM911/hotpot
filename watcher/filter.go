package watcher

import (
	"os"

	"github.com/KM911/fish/fs"
)

func Folders() []string {
	var folders []string
	files, _ := os.ReadDir(fs.WorkingDirectory)
	folders = append(folders, fs.WorkingDirectory)
	for _, file := range files {
		if file.IsDir() {
			if _, ok = IgnoreFolders[file.Name()]; !ok {
				folders = append(folders, fs.WorkingDirectory+string(os.PathSeparator)+file.Name())
			}
		}
	}
	return folders
}
