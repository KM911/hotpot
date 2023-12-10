package watcher

import (
	"os"

	"github.com/KM911/hotpot/util"
)

func Folders() []string {
	var folders []string
	files, _ := os.ReadDir(util.ExecuteDirectory)
	folders = append(folders, util.ExecuteDirectory)
	for _, file := range files {
		if file.IsDir() {
			if _, ok = IgnoreFolders[file.Name()]; !ok {
				folders = append(folders, util.ExecuteDirectory+string(os.PathSeparator)+file.Name())
			}
		}
	}
	return folders
}
