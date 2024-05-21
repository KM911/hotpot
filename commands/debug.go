package commands

import (
	"fmt"

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

	fmt.Println(c.Args().Slice())
	watcher.WatchHook(nil)
	return nil
}
