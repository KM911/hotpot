package watcher

import (
	"fmt"
	"github.com/KM911/hotpot/config"
	"log"

	"github.com/KM911/hotpot/commands"
)

func StartWatch() {
	defer watcher.Close()
	fmt.Println("Start Watch")
	fmt.Print("Watch file type：")
	for _, v := range config.UserToml.WatchFiles {
		fmt.Print(v, " ")
	}
	fmt.Println("\nExecute command:", config.UserToml.Command)
	commands.Start()
	for {
		select {
		// this will ignore .sh file ? why
		case event = <-watcher.Events:
			if config.UserToml.ShowEvent {
				fmt.Println("event:", event.Name, event.Op)
			}
			Debounce(func() {
				EventHandle(event)
			})
		case err, ok = <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("error:", err)
		}
	}
}
