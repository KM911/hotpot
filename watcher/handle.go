package watcher

import (
	"path/filepath"

	"github.com/KM911/hotpot/commands"

	"github.com/fsnotify/fsnotify"
)

func EventHandle(event fsnotify.Event) {
	ext := filepath.Ext(event.Name)
	if ext == "" {
		return
	}
	if _, ok := WatchFiles[ext[1:]]; ok {
		commands.Stop()
		println("\033[H\033[2J")
		commands.Start()
	}
}
