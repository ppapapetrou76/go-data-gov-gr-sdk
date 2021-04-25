package traffic

import (
	"errors"
	"testing"
	"time"

	"github.com/ppapapetrou76/go-testing/assert"

	"github.com/ppapapetrou76/go-data-gov-gr-sdk/api"
	"github.com/ppapapetrou76/go-data-gov-gr-sdk/internal"
)

const sampleData = `[
  {
    "deviceid":"321",
    "countedcars":1000,
    "appprocesstime":"2021-04-25T05:00:00Z",
    "road_name":"ΚΑΤΕΧΑΚΗ",
    "road_info":"ΚΑΤΕΧΑΚΗ ΔΕΥΤΕΡΗ ΕΞΟΔΟΣ",
    "average_speed":25.123
  },
  {
    "deviceid":"123",
    "countedcars":1500,
    "appprocesstime":"2021-04-25T05:00:00Z",
    "road_name":"ΚΑΤΕΧΑΚΗ",
    "road_info":"ΚΑΤΕΧΑΚΗ ΕΞΟΔΟΣ",
    "average_speed":65.987
  }
]`

func expectedList() List {
	return List{
		{
			DeviceID:     "321",
			CountedCars:  1000,
			ProcessTime:  time.Date(2021, 4, 25, 5, 0, 0, 0, time.UTC),
			RoadName:     "ΚΑΤΕΧΑΚΗ",
			RoadInfo:     "ΚΑΤΕΧΑΚΗ ΔΕΥΤΕΡΗ ΕΞΟΔΟΣ",
			AverageSpeed: 25.123,
		},
		{
			DeviceID:     "123",
			CountedCars:  1500,
			ProcessTime:  time.Date(2021, 4, 25, 5, 0, 0, 0, time.UTC),
			RoadName:     "ΚΑΤΕΧΑΚΗ",
			RoadInfo:     "ΚΑΤΕΧΑΚΗ ΕΞΟΔΟΣ",
			AverageSpeed: 65.987,
		},
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		name        string
		params      *api.GetParams
		client      *api.Client
		expected    List
		expectedErr error
	}{
		{
			name:        "should fail if API returns an error",
			client:      internal.NewCommonMockClientError("road_traffic_attica", "cannot fetch data"),
			params:      api.NewDefaultGetParams(),
			expectedErr: errors.New("get traffic data: cannot fetch data"),
		},
		{
			name:     "should succeed when API call succeeds",
			client:   internal.NewCommonMockClientSuccess("road_traffic_attica", sampleData),
			params:   api.NewDefaultGetParams(),
			expected: expectedList(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := Get(tt.client, tt.params)
			assert.ThatError(t, err).IsSameAs(tt.expectedErr)
			assert.That(t, actual).IsEqualTo(tt.expected)
		})
	}
}
