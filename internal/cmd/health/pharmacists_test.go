package cmdhealth

import (
	"testing"

	"github.com/ppapapetrou76/go-testing/assert"

	cmdglobal "github.com/ppapapetrou76/go-data-gov-gr-sdk/internal/cmd/global"
)

func Test_pharmacistCmd(t *testing.T) {
	cmd := pharmacistCmd()

	assert.That(t, cmd.Name).IsEqualTo("pharmacist")
	assert.That(t, cmd.Usage).IsEqualTo("Shows pharmacists statistics")
	assert.ThatSlice(t, cmd.Flags).
		HasSize(4).
		Contains(cmdglobal.CommonFlags()).
		Contains(cmdglobal.YearRangeFlags())
}
