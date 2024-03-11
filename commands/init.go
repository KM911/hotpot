package commands

import (
	"github.com/urfave/cli/v2"
)

var (
	Init = cli.Command{
		Name:    "init",
		Usage:   "Init config file",
		Aliases: []string{"i"},
		Action:  InitAction,
	}
)

func InitAction(c *cli.Context) error {
	// show info
	return nil
}
