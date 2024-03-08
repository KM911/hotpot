package watcher

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/KM911/hotpot/commands"
	"github.com/KM911/hotpot/config"
	"github.com/KM911/hotpot/lib/format"
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
			//TODO change the function in init
			// step1 : show event
			// step2 : reload config file
			// step3 : handle event
			if config.UserToml.ShowEvent {
				format.InfoMessage(event.Op.String(), event.Name)
			}
			if strings.Contains(event.Name, config.TomlFile) {
				fmt.Println("reload config file")
				config.LoadToml()
				ProcessWatchEnvironment()
			} else {
				// problem is that

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
