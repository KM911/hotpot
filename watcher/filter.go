package watcher

import (
	"hotpot/util"
	"os"
)

func Folders() []string {
	var folders []string
	files, _ := os.ReadDir(util.RootPath)
	folders = append(folders, util.RootPath)
	for _, file := range files {
		if file.IsDir() {
			if _, ok = IgnoreFolders[file.Name()]; !ok {
				folders = append(folders, util.RootPath+string(os.PathSeparator)+file.Name())
			}
		}
	}
	return folders
}
