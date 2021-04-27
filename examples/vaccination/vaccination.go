package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ppapapetrou76/go-data-gov-gr-sdk/api"

	"github.com/ppapapetrou76/go-data-gov-gr-sdk/pkg/vaccination"
)

func main() {
	// Fetches the vaccination data for the last 6 days for all areas
	client := api.NewClient("<YOUR_API_TOKEN_HERE>")
	data, err := vaccination.Get(client,
		api.NewDefaultGetParams(api.SetDateFrom(time.Now().Add(-time.Hour*24*5))),
	)
	if err != nil {
		panic(err)
	}
	// Filter by a specific region
	for _, d := range data.FilterByArea("ΘΕΣΣΑΛΟΝΙΚΗΣ") {
		fmt.Fprintf(os.Stdout, "Area:%s, Vaccinations on %v:%d\n", d.Area, d.ReferenceDate, d.DayTotal)
	}
}
