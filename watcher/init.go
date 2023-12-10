package watcher

import (
	"github.com/KM911/hotpot/format"
	"log"
	"time"

	"github.com/KM911/hotpot/util"

	"github.com/KM911/hotpot/config"

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
	//执行的先后顺序是什么? 这里才是最关键的
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
	// util.DrawBlock("watch folder", watchFolder)
	format.BlockMessage("watch folder", watchFolder)
}
