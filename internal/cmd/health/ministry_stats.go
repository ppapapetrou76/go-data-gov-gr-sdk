package cmdhealth

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/ppapapetrou76/go-data-gov-gr-sdk/api"
	"github.com/ppapapetrou76/go-data-gov-gr-sdk/internal"
	cmdglobal "github.com/ppapapetrou76/go-data-gov-gr-sdk/internal/cmd/global"
	"github.com/ppapapetrou76/go-data-gov-gr-sdk/internal/formatter"
	"github.com/ppapapetrou76/go-data-gov-gr-sdk/pkg/health/ministrystats"
)

func ministryStatsCmd() *cli.Command {
	flags := cmdglobal.CommonFlags()
	flags = append(flags, cmdglobal.YearRangeFlags()...)
	flags = append(flags, &cli.StringFlag{
		Name:    "category",
		Aliases: []string{"c"},
		Usage:   "Used to specify which category of data should be fetched (valid options: pharmacists|pharmacies|doctors|dentists",
	})

	return &cli.Command{
		Name:  "ministry-stats",
		Usage: "Shows ministry statistics",
		Flags: flags,
		Action: func(context *cli.Context) error {
			client := api.NewClient(context.String("auth-token"),
				api.SetHTTPClient(internal.Get().HTTPClient),
			)
			data, err := ministrystats.Get(client, ministrystats.GetParams{
				GetParams: api.NewDefaultGetParams(),
				Category:  ministrystats.Category(context.String("category")),
			})
			if err != nil {
				return fmt.Errorf("ministry statistics:%w", err)
			}
			data = data.FilterByYearRange(context.Int("year-from"), context.Int("year-to"))
			if err := formatter.New(context.App.Writer, context.String("output")).Format(data); err != nil {
				return fmt.Errorf("ministry statistics:%w", err)
			}

			return nil
		},
	}
}
