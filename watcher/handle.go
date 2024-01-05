package watcher

import (
	"github.com/KM911/hotpot/config"
	"path/filepath"

	"github.com/KM911/hotpot/commands"

	"github.com/fsnotify/fsnotify"
)

var (
	EventHandle func(event fsnotify.Event)
)

//func EventHandle(event fsnotify.Event) {
//
//	ext := filepath.Ext(event.Name)
//	if ext == "" {
//		return
//	}
//	if _, ok := WatchFiles[ext[1:]]; ok {
//		commands.Stop()
//		println("\033[H\033[2J")
//		commands.Start()
//	}
//}

func init() {
	if len(config.UserToml.WatchFiles) == 1 && config.UserToml.WatchFiles[0] == "*" {
		EventHandle = EventHandleAll
	} else {
		EventHandle = EventHandleWithFileExtension
	}
}

func EventHandleWithFileExtension(event fsnotify.Event) {
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

func EventHandleAll(event fsnotify.Event) {
	commands.Stop()
	println("\033[H\033[2J")
	commands.Start()
}
