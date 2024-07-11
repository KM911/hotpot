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
		Clear()
		Start()
	}
}

func EventHandleAll(event fsnotify.Event) {
	Stop()
	Clear()
	Start()
}
func Clear() {
	println("\033[H\033[2J")
}
