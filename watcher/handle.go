package watcher

import (
	"hotpot/commands"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

func EventHandle(event fsnotify.Event) {
	ext := filepath.Ext(event.Name)
	// fmt.Println("change file extension is ", ext)
	if ext == "" {
		return
	}
	if _, watch := WatchFiles[ext[1:]]; watch {
		commands.Stop()
		commands.Start()
	}
}
