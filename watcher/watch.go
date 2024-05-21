package watcher

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/KM911/hotpot/config"
	"github.com/KM911/hotpot/lib/format"
	"github.com/fsnotify/fsnotify"
)

func StartWatch() {
	defer watcher.Close()

	format.NoteMessage("Ignores", strings.Join(config.UserToml.IgnoreFolders, ","))
	format.InfoMessage("File Type", strings.Join(config.UserToml.WatchFiles, ","))
	// format.NoteMessage("Command", config.UserToml.Command)
	format.InfoMessage("Delay", strconv.Itoa(config.UserToml.Delay))
	println("------------------------------------------")
	Start()
	for {
		select {
		case event = <-watcher.Events:
			if config.UserToml.ShowEvent {
				format.InfoMessage(event.Op.String(), event.Name)
			}
			if strings.Contains(event.Name, config.TomlFile) {
				fmt.Println("reload config file")
				config.LoadToml()
				ProcessWatchEnvironment()
			} else {
				Debounce(func() {
					EventHandle(event)
				})
			}
		case err, ok = <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("error:", err)
		}
	}
}

func WatchHook(f func(event fsnotify.Event)) {

	defer watcher.Close()

	format.NoteMessage("Ignores", strings.Join(config.UserToml.IgnoreFolders, ","))
	format.InfoMessage("File Type", strings.Join(config.UserToml.WatchFiles, ","))
	format.InfoMessage("Delay", strconv.Itoa(config.UserToml.Delay))
	println("------------------------------------------")
	Start()
	for {
		select {
		case event = <-watcher.Events:
			if config.UserToml.ShowEvent {
				format.InfoMessage(event.Op.String(), event.Name)
			}
			// if strings.Contains(event.Name, config.TomlFile) {
			// 	fmt.Println("reload config file")
			// 	config.LoadToml()
			// 	ProcessWatchEnvironment()
			// } else {
			// 	Debounce(func() {
			// 		EventHandle(event)
			// 	})
			// }
		case err, ok = <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("error:", err)
		}
	}
}
