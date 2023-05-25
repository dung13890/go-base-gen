package cmd

import (
	"log"

	"github.com/urfave/cli/v2"
)

// NewProject is a function to create new project command
func NewProject() *cli.Command {
	return &cli.Command{
		Name:  "project",
		Usage: "Generate base code for go project use clean architecture",
		Action: func(ctx *cli.Context) error {
			log.Println("project creating: ", ctx.Args().First())
			return nil
		},
	}
}
