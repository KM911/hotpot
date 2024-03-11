package commands

import (
	"github.com/KM911/hotpot/watcher"
	"github.com/urfave/cli/v2"
)

var (
	Watch = cli.Command{
		Name:    "watch",
		Usage:   "Start watch file change and execute command",
		Aliases: []string{"s", "w"},
		Action:  WatchAction,
	}
)

func WatchAction(c *cli.Context) error {
	watcher.StartWatch()
	return nil
}
