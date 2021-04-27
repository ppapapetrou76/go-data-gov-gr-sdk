package cmdhealth

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/ppapapetrou76/go-data-gov-gr-sdk/api"
	cmdglobal "github.com/ppapapetrou76/go-data-gov-gr-sdk/internal/cmd/global"
	"github.com/ppapapetrou76/go-data-gov-gr-sdk/internal/formatter"
	"github.com/ppapapetrou76/go-data-gov-gr-sdk/pkg/vaccination"
)

func vaccinationCmd() *cli.Command {
	flags := append(cmdglobal.Flags(),
		&cli.StringFlag{
			Name:  "area-name",
			Usage: "Filters the results based on a given area name",
		})

	return &cli.Command{
		Name:  "vaccination",
		Usage: "Shows vaccination statistics",
		Flags: flags,
		Action: func(context *cli.Context) error {
			client := api.NewClient(context.String("auth-token"))
			data, err := vaccination.Get(client, api.NewDefaultGetParams(
				api.SetDateFrom(*context.Timestamp("date-from")),
				api.SetDateTo(*context.Timestamp("date-to")),
			))
			if err != nil {
				return fmt.Errorf("vaccination statistics:%w", err)
			}
			data = data.FilterByArea(context.String("area-name"))

			if err := formatter.New(os.Stdout, context.String("output")).Format(data); err != nil {
				return fmt.Errorf("vaccination statistics:%w", nil)
			}

			return nil
		},
	}
}
