package watcher

import (
	"os"

	"github.com/KM911/hotpot/lib/util"
)

func Folders() []string {
	var folders []string
	files, _ := os.ReadDir(util.WorkingDirectory)
	folders = append(folders, util.WorkingDirectory)
	for _, file := range files {
		if file.IsDir() {
			if _, ok = IgnoreFolders[file.Name()]; !ok {
				folders = append(folders, util.WorkingDirectory+string(os.PathSeparator)+file.Name())
			}
		}
	}
	return folders
}
