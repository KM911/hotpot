package app

import (
	"log"
	"os"

	"github.com/KM911/hotpot/commands"

	"github.com/urfave/cli/v2"
)

var ()

func Template() {
	app := &cli.App{
		// TODO set basic info for app
		Name:     "hotpot",
		Usage:    "Create hard link for project",
		Commands: commands.Subcommands,
		Action: func(_cli_context *cli.Context) error {
			cli.ShowAppHelpAndExit(_cli_context, 0)
			return nil
		},
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "KM911",
				Email: "2547715095@qq.com",
			},
		},
		Suggest:  true,
		HideHelp: true,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func NewApp(_name, _usage string) {
	app := &cli.App{
		Name:     _name,
		Usage:    _usage,
		Commands: commands.Subcommands,
		Action: func(_cli_context *cli.Context) error {
			cli.ShowAppHelpAndExit(_cli_context, 0)
			return nil
		},
		// Authors: ,

		Suggest:  true,
		HideHelp: true,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
