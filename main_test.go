package main

import (
	"testing"

	"github.com/ppapapetrou76/go-testing/assert"
	"github.com/urfave/cli/v2"

	"github.com/ppapapetrou76/go-data-gov-gr-sdk/internal/cmd"
)

func Test_cliApp(t *testing.T) {
	actual := cliApp()
	expected := &cli.App{
		Name:  "ggd-cli",
		Usage: "The opensource CLI of open data available in data.gov.gr by the Greek Government",
		Authors: []*cli.Author{
			{Name: "Patroklos Papapetrou", Email: "ppapapetrou76 at gmail dot com"},
		},
		Commands: cmd.Commands(),
	}

	assert.ThatStruct(t, actual).
		ExcludingFields("Commands", "didSetup", "categories", "flagCategories", "rootCommand").
		IsEqualTo(expected)
}
