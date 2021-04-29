package cmdhealth

import (
	"errors"
	"flag"
	"fmt"
	"testing"
	"time"

	"github.com/ppapapetrou76/go-testing/assert"
	"github.com/urfave/cli/v2"

	"github.com/ppapapetrou76/go-data-gov-gr-sdk/api"
	"github.com/ppapapetrou76/go-data-gov-gr-sdk/internal"
	cmdglobal "github.com/ppapapetrou76/go-data-gov-gr-sdk/internal/cmd/global"
)

func Test_pharmacistCmd(t *testing.T) {
	cmd := ministryStatsCmd()

	assert.That(t, cmd.Name).IsEqualTo("ministry-stats")
	assert.That(t, cmd.Usage).IsEqualTo("Shows ministry statistics")
	assert.ThatSlice(t, cmd.Flags).
		HasSize(5).
		Contains(cmdglobal.CommonFlags()).
		Contains(cmdglobal.YearRangeFlags()).
		Contains(&cli.StringFlag{
			Name:    "category",
			Aliases: []string{"c"},
			Usage:   "Used to specify which category of data should be fetched (valid options: pharmacists|pharmacies|doctors|dentists",
		})
}

func Test_pharmacistCmd_Action(t *testing.T) {
	sampleData := `[
  {
    "year": 2018,
    "quarter": "Q3",
    "active": 10616,
    "entrants": 30,
    "exits": 53
  },
  {
    "year": 2018,
    "quarter": "Q4",
    "active": 10587,
    "entrants": 40,
    "exits": 69
  }
]
`
	internal.Instance(nil)
	tests := []struct {
		name        string
		expectedErr error
		httpClient  api.HTTPClient
		flagSet     *flag.FlagSet
	}{
		{
			name:        "should error if getting ministrystats data fails",
			flagSet:     internal.NewFlagSet(),
			expectedErr: errors.New("ministry statistics:ministry stats get validation: category  is not recognized. valid categories are: [pharmacists pharmacies doctors dentists]"),
		},
		{
			name: "should succeed",
			httpClient: api.NewMockHTTPClient(api.MockRequest{
				Path:         "minhealth_pharmacists",
				Query:        fmt.Sprintf("date_to=%s", time.Now().Format("2006-01-02")),
				ResponseBody: api.NewMockBody(sampleData),
			}),
			flagSet: flagSetWithCategoryDefined("pharmacists"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := ministryStatsCmd()
			internal.Get().HTTPClient = tt.httpClient
			err := cmd.Action(internal.NewCLIContext(tt.flagSet))
			assert.ThatError(t, err).IsSameAs(tt.expectedErr)
		})
	}
}

func flagSetWithCategoryDefined(category string) *flag.FlagSet {
	flagSet := internal.NewFlagSet()
	flagSet.String("category", category, "")

	return flagSet
}
