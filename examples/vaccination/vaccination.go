package main

import (
	"fmt"
	vaccination2 "github.com/ppapapetrou76/go-data-gov-gr-sdk/pkg/health/vaccination"
	"os"
	"time"

	"github.com/ppapapetrou76/go-data-gov-gr-sdk/api"
)

func main() {
	// Fetches the vaccination data for the last 6 days for all areas
	client := api.NewClient("<YOUR_API_TOKEN_HERE>")
	data, err := vaccination2.Get(client,
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
