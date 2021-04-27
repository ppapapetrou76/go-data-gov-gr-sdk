package cmd

import (
	"github.com/urfave/cli/v2"

	cmdhealth "github.com/ppapapetrou76/go-data-gov-gr-sdk/internal/cmd/health"
)

// Commands returns the list of top-level CLI commands.
func Commands() []*cli.Command {
	return []*cli.Command{
		cmdhealth.Command(),
	}
}
