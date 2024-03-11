package commands

import (
	"fmt"
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
	config.UserToml.Command = strings.Join(c.Args().Slice(), " ")
	fmt.Println(c.Args().Slice())
	watcher.StartWatch()
	return nil
}
