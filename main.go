package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/ppapapetrou76/go-data-gov-gr-sdk/internal/cmd"
)

func main() {
	app := &cli.App{
		Name:  "ggd-cli",
		Usage: "The opensource CLI of open data available in data.gov.gr by the Greek Government",
		Authors: []*cli.Author{
			{Name: "Patroklos Papapetrou", Email: "ppapapetrou76 at gmail dot com"},
		},
		Commands: cmd.Commands(),
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
