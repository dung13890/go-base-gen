package main

import (
	"log"
	"os"

	"github.com/dung13890/go-base-gen/cmd"
	"github.com/urfave/cli/v2"
)

var (
	version string = "v1.0.9"
)

func main() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "print only the version",
	}

	project := cmd.NewProject()
	domain := cmd.NewDomain()

	app := &cli.App{
		Name:    "go-base-gen",
		Suggest: true,
		Version: version,
		Usage:   "Use this tool to generate base code",
		Commands: []*cli.Command{
			project,
			domain,
		},
		Action: func(ctx *cli.Context) error {
			if err := cli.ShowAppHelp(ctx); err != nil {
				return err
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
