package watcher

import (
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

var (
	EventHandle func(event fsnotify.Event)
	HookSocket  func()
)

func EventHandleWithFileExtension(event fsnotify.Event) {
	ext := filepath.Ext(event.Name)
	if ext == "" {
		return
	}
	if _, ok := WatchFiles[ext[1:]]; ok {
		Stop()
		println("\033[H\033[2J")
		Start()
	}
}

func EventHandleAll(event fsnotify.Event) {
	Stop()
	println("\033[H\033[2J")
	Start()
}
