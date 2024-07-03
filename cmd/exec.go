package cmd

import (
	"strings"

	"github.com/KM911/hotpot/config"
	"github.com/KM911/hotpot/watcher"
	"github.com/urfave/cli/v2"
)

var (
	Exec = cli.Command{
		Name:    "exec",
		Usage:   "Use argv as command and watch",
		Aliases: []string{"e"},
		Action:  ExecAction,
	}
)

func ExecAction(c *cli.Context) error {
	config.UserToml.ExecuteCommand = strings.Join(c.Args().Slice(), " ")
	watcher.ProcessWatchEnvironment()
	watcher.StartWatch()
	return nil
}
