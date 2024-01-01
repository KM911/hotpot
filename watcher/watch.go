package watcher

import (
	"log"
	"strconv"
	"strings"

	"github.com/KM911/hotpot/commands"
	"github.com/KM911/hotpot/config"
	"github.com/KM911/hotpot/format"
)

func StartWatch() {
	defer watcher.Close()
	//format.DrawBlock("watch file type", config.UserToml.WatchFiles)
	//format.DrawBlock("ignores", config.UserToml.IgnoreFolders)
	format.NoteMessage("Ignores", strings.Join(config.UserToml.IgnoreFolders, ","))
	format.InfoMessage("File Type", strings.Join(config.UserToml.WatchFiles, ","))
	format.NoteMessage("Command", config.UserToml.Command)
	format.InfoMessage("Delay", strconv.Itoa(config.UserToml.Delay))

	println("------------------------------------------")

	commands.Start()

	//TODO : decrease the if
	for {
		select {
		case event = <-watcher.Events:
			if config.UserToml.ShowEvent {
				format.InfoMessage(event.Op.String(), event.Name)
			}
			if strings.Contains(event.Name, config.TomlFile) {
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
