package cmdglobal

import (
	"time"

	"github.com/urfave/cli/v2"

	"github.com/ppapapetrou76/go-data-gov-gr-sdk/pkg/util/times"
)

// CommonFlags returns the global flags used for all commands/subcommands.
func CommonFlags() []cli.Flag {
	return []cli.Flag{
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
}

// DateRangeFlags returns the commonly used date range flags.
func DateRangeFlags() []cli.Flag {
	return []cli.Flag{
		&cli.TimestampFlag{
			Name:        "date-from",
			Usage:       "Used to fetch data after a given date",
			Layout:      "2006-01-02",
			Value:       cli.NewTimestamp(time.Now().Add(-times.Week)),
			DefaultText: "one week before",
		},
		&cli.TimestampFlag{
			Name:        "date-to",
			Usage:       "Used to fetch data before a given date",
			Layout:      "2006-01-02",
			Value:       cli.NewTimestamp(time.Now()),
			DefaultText: "today",
		},
	}
}

// YearRangeFlags returns the commonly used year range flags.
func YearRangeFlags() []cli.Flag {
	return []cli.Flag{
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
}
