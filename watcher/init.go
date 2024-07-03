package watcher

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/KM911/fish/adt"
	"github.com/KM911/fish/format"
	"github.com/KM911/fish/system"
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
	cmd      *exec.Cmd
	hookCmds = make(chan *exec.Cmd, 10)
)

func ProcessWatchEnvironment() {
	WatchFiles = map[string]struct{}{}
	IgnoreFolders = map[string]struct{}{}
	adt.SetAppend(WatchFiles, config.UserToml.WatchFiles)
	adt.SetAppend(IgnoreFolders, config.UserToml.IgnoreFolders)

	Debounce = system.NewDebounce(time.Duration(config.UserToml.Delay) * time.Millisecond)
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

	if len(config.UserToml.WatchFiles) == 1 && config.UserToml.WatchFiles[0] == "*" {
		EventHandle = EventHandleAll
	} else {
		EventHandle = EventHandleWithFileExtension
	}

	format.NoteMessage("Ignores", strings.Join(config.UserToml.IgnoreFolders, ","))
	format.InfoMessage("File Type", strings.Join(config.UserToml.WatchFiles, ","))
	format.NoteMessage("Execute Command", config.UserToml.ExecuteCommand)
	format.InfoMessage("Delay", strconv.Itoa(config.UserToml.Delay))
	format.InfoMessage("ShowEvent", strconv.FormatBool(config.UserToml.ShowEvent))

}
