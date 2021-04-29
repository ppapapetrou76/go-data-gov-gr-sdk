package cmdglobal

import (
	"testing"
	"time"

	"github.com/ppapapetrou76/go-testing/assert"
	"github.com/urfave/cli/v2"

	"github.com/ppapapetrou76/go-data-gov-gr-sdk/pkg/util/times"
)

func TestCommonFlags(t *testing.T) {
	actualFlags := CommonFlags()
	expectedFlags := []cli.Flag{
		&cli.StringFlag{
			Name:     "auth-token",
			Usage:    "Used for API authentication",
			EnvVars:  []string{"DATA_GOV_GR_API_TOKEN"},
			Required: true,
		},
		&cli.StringFlag{
			Name:        "output",
			Usage:       "Output format [json|yaml|text]",
			DefaultText: "json",
		},
	}

	assert.That(t, actualFlags).IsEqualTo(expectedFlags)
}

func TestYearRangeFlags(t *testing.T) {
	actualFlags := YearRangeFlags()
	expectedFlags := []cli.Flag{
		&cli.IntFlag{
			Name:  "year-from",
			Usage: "Used to fetch data after a given year",
		},
		&cli.IntFlag{
			Name:        "year-to",
			Usage:       "Used to fetch data before a given year",
			Value:       time.Now().Year(),
			DefaultText: "current year",
		},
	}

	assert.That(t, actualFlags).IsEqualTo(expectedFlags)
}

func TestDateRangeFlags(t *testing.T) {
	actualFlags := DateRangeFlags()
	now := time.Now()
	roundedNow := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	expectedFlags := []cli.Flag{
		&cli.TimestampFlag{
			Name:        "date-from",
			Usage:       "Used to fetch data after a given date",
			Layout:      "2006-01-02",
			Value:       cli.NewTimestamp(roundedNow.Add(-times.Week)),
			DefaultText: "one week before",
		},
		&cli.TimestampFlag{
			Name:        "date-to",
			Usage:       "Used to fetch data before a given date",
			Layout:      "2006-01-02",
			Value:       cli.NewTimestamp(roundedNow),
			DefaultText: "today",
		},
	}

	assert.That(t, actualFlags).IsEqualTo(expectedFlags)
}
