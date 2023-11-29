package watcher

import (
	"path/filepath"

	"github.com/KM911/hotpot/commands"

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
