package commands

import (
	"github.com/KM911/hotpot/config"
	"github.com/KM911/hotpot/watcher"
	"github.com/urfave/cli/v2"
)

var (
	Debug = cli.Command{
		Name:    "debug",
		Usage:   "show file change event",
		Aliases: []string{"d"},
		Action:  DebugAction,
	}
)

func DebugAction(c *cli.Context) error {
	watcher.ProcessWatchEnvironment()
	config.UserToml.ExecuteCommand = ""
	config.UserToml.ShowEvent = true
	watcher.StartWatch()
	return nil
}
