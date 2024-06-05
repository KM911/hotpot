package commands

import (
	"github.com/KM911/hotpot/config"
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
	config.LoadToml()
	return nil
}
