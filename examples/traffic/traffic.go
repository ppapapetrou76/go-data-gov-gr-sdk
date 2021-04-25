package traffic

import (
	"fmt"
	"os"
	"time"

	"github.com/ppapapetrou76/go-data-gov-gr-sdk/api"
	"github.com/ppapapetrou76/go-data-gov-gr-sdk/pkg/traffic"
)

func main() {
	// Fetches the traffic data in Attica for the last 2 hours for all areas
	client := api.NewClient("<YOUR_API_TOKEN_HERE>")
	trafficData, err := traffic.Get(client,
		api.NewDefaultGetParams(api.SetDateFrom(time.Now().Add(-time.Hour * 2))))
	if err != nil {
		panic(err)
	}
	// Filter by a specific road
	for _, d := range trafficData.FilterByRoadName("Λ. ΜΕΣΟΓΕΙΩΝ") {
		fmt.Fprintf(os.Stdout, "Area:%s, Counted cars on %v(%v):%d with average speed of %f\n",
			d.RoadName, d.RoadInfo, d.ProcessTime, d.CountedCars, d.AverageSpeed)
	}
}
