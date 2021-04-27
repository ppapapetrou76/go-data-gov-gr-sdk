package cmdhealth

import "github.com/urfave/cli/v2"

// Command return the `health` top-level command.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "health",
		Usage: "the top-level command for all health-related sub-commands",
		Subcommands: []*cli.Command{
			vaccinationCmd(),
			foodInspectionCmd(),
		},
	}
}
