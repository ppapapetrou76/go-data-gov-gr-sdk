package cmdhealth

import "github.com/urfave/cli/v2"

// Command return the `health` top-level command.
func Command() *cli.Command {
	return &cli.Command{
		Name:         "health",
		Aliases:      nil,
		Usage:        "the top-level command for all health-related sub-commands",
		UsageText:    "",
		Description:  "",
		ArgsUsage:    "",
		Category:     "",
		BashComplete: nil,
		Before:       nil,
		After:        nil,
		Action:       nil,
		OnUsageError: nil,
		Subcommands: []*cli.Command{
			vaccinationCmd(),
		},
		Flags:                  nil,
		SkipFlagParsing:        false,
		HideHelp:               false,
		HideHelpCommand:        false,
		Hidden:                 false,
		UseShortOptionHandling: false,
		HelpName:               "",
		CustomHelpTemplate:     "",
	}
}
