package watcher

import (
	"log"
	"time"

	"github.com/KM911/hotpot/config"
	"github.com/KM911/hotpot/lib/format"
	"github.com/KM911/hotpot/lib/util"
	"github.com/fsnotify/fsnotify"
)

var (
	IgnoreFolders = map[string]struct{}{}
	WatchFiles    = map[string]struct{}{}

	watcher  *fsnotify.Watcher
	err      error
	ok       bool
	Debounce func(func())
	event    fsnotify.Event
)

func init() {
	ProcessWatchEnvironment()
}
func ProcessWatchEnvironment() {
	WatchFiles = map[string]struct{}{}
	IgnoreFolders = map[string]struct{}{}
	util.SetAppend(WatchFiles, config.UserToml.WatchFiles)
	util.SetAppend(IgnoreFolders, config.UserToml.IgnoreFolders)

	Debounce = util.NewDebounce(time.Duration(config.UserToml.Delay) * time.Millisecond)
	watcher, err = fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	watchFolder := []string{}
	for _, folder := range Folders() {
		if _, ok := IgnoreFolders[folder]; ok {
			continue
		}
		watchFolder = append(watchFolder, folder)
		err = watcher.Add(folder)
		if err != nil {
			log.Fatal(err)
		}
	}

	format.BlockMessage("watch folder", watchFolder)
}
