package cmdhealth

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/ppapapetrou76/go-data-gov-gr-sdk/api"
	cmdglobal "github.com/ppapapetrou76/go-data-gov-gr-sdk/internal/cmd/global"
	"github.com/ppapapetrou76/go-data-gov-gr-sdk/internal/formatter"
	"github.com/ppapapetrou76/go-data-gov-gr-sdk/pkg/foodinspection"
)

func foodInspectionCmd() *cli.Command {
	flags := cmdglobal.CommonFlags()
	flags = append(flags, cmdglobal.YearRangeFlags()...)

	return &cli.Command{
		Name:  "food-inspection",
		Usage: "Shows food inspection statistics",
		Flags: flags,
		Action: func(context *cli.Context) error {
			client := api.NewClient(context.String("auth-token"))
			data, err := foodinspection.Get(client, api.NewDefaultGetParams())
			if err != nil {
				return fmt.Errorf("food inspection statistics:%w", err)
			}
			data = data.FilterByYearRange(context.Int("year-from"), context.Int("year-to"))
			if err := formatter.New(os.Stdout, context.String("output")).Format(data); err != nil {
				return fmt.Errorf("vaccination statistics:%w", err)
			}

			return nil
		},
	}
}
