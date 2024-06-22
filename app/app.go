package app

import (
	"log"
	"os"

	"github.com/KM911/hotpot/commands"

	"github.com/urfave/cli/v2"
)

var ()

func NewCliApp() *cli.App {
	return &cli.App{
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
}

func Template() {
	app := NewCliApp()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func NewApp(_name, _usage string) {
	app := NewCliApp()
	app.Name = _name
	app.Usage = _usage
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
