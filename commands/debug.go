package commands

import (
	"fmt"
	"strings"

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
	config.UserToml.Command = strings.Join(c.Args().Slice(), " ")
	fmt.Println(c.Args().Slice())
	watcher.WatchHook(nil)
	return nil
}
