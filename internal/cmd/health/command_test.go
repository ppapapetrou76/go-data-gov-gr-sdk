package cmdhealth

import (
	"testing"

	"github.com/ppapapetrou76/go-testing/assert"
	"github.com/urfave/cli/v2"
)

func TestCommand(t *testing.T) {
	actualCommands := Command()
	expectedCommands := &cli.Command{
		Name:  "health",
		Usage: "the top-level command for all health-related sub-commands",
		Subcommands: []*cli.Command{
			vaccinationCmd(),
			foodInspectionCmd(),
			pharmacistCmd(),
		},
	}
	// We are excluding sub-commands because each one is containing an action func and comparison of funcs always
	// returns false
	assert.ThatStruct(t, actualCommands).ExcludingFields("Subcommands").IsEqualTo(expectedCommands)

	assert.ThatSlice(t, actualCommands.Subcommands).HasSize(3)
	for i := range actualCommands.Subcommands {
		assert.ThatStruct(t, actualCommands.Subcommands[i]).ExcludingFields("Action").IsEqualTo(expectedCommands.Subcommands[i])
	}
}
