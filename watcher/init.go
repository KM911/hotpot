package watcher

import (
	"fmt"
	"hotpot/config"
	"hotpot/util"
	"log"
	"time"

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
	util.SetAppend(WatchFiles, config.UserToml.WatchFiles)
	util.SetAppend(IgnoreFolders, config.UserToml.IgnoreFolders)

	Debounce = util.NewDebounce(time.Duration(config.UserToml.Delay) * time.Millisecond)

	watcher, err = fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	for _, folder := range Folders() {
		// filter ignore folders
		if _, ok := IgnoreFolders[folder]; ok {
			continue
		}
		fmt.Println("watch folder ：", folder)
		err = watcher.Add(folder)
		if err != nil {
			log.Fatal(err)
		}
	}

}
