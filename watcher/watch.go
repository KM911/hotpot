package watcher

import (
	"log"
	"strings"

	"github.com/KM911/hotpot/config"
	"github.com/KM911/oslib/format"
)

func StartWatch() {
	defer watcher.Close()
	if config.HookEnable {
		ConnectSocket()
		SocketPing = SocketPingAction
	} else {
		SocketPing = SocketPingEmpty
	}
	Start()
	for {
		select {
		case event = <-watcher.Events:
			if config.UserToml.ShowEvent {
				format.InfoMessage(event.Op.String(), event.Name)
			}
			if strings.Contains(event.Name, config.TomlFile) {
				format.Info("Reload Config File")
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
