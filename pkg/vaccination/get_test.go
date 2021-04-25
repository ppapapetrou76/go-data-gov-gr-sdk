package vaccination

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/ppapapetrou76/go-testing/assert"

	"github.com/ppapapetrou76/go-data-gov-gr-sdk/api"
)

const sampleData = `[
  {
    "area":"ΑΙΤΩΛΟΑΚΑΡΝΑΝΙΑΣ",
    "areaid":701,
    "dailydose1":665,
    "dailydose2":171,
    "daydiff":118,
    "daytotal":836,
    "referencedate":"2021-04-22T00:00:00",
    "totaldistinctpersons":27193,
    "totaldose1":27193,
    "totaldose2":7111,
    "totalvaccinations":34304
  },
  {
    "area":"ΑΝΑΤΟΛΙΚΗΣ ΑΤΤΙΚΗΣ",
    "areaid":901,
    "dailydose1":1054,
    "dailydose2":107,
    "daydiff":-39,
    "daytotal":1161,
    "referencedate":"2021-04-22T00:00:00",
    "totaldistinctpersons":51872,
    "totaldose1":51872,
    "totaldose2":17531,
    "totalvaccinations":69403
  }
]`

func expectedList() List {
	return List{
		{
			Area:                 "ΑΙΤΩΛΟΑΚΑΡΝΑΝΙΑΣ",
			AreaID:               701,
			DailyShot1:           665,
			DailyShot2:           171,
			DayDiff:              118,
			DayTotal:             836,
			ReferenceDate:        "2021-04-22T00:00:00",
			TotalDistinctPersons: 27193,
			TotalShot1:           27193,
			TotalShot2:           7111,
			TotalVaccinations:    34304,
		},
		{
			Area:                 "ΑΝΑΤΟΛΙΚΗΣ ΑΤΤΙΚΗΣ",
			AreaID:               901,
			DailyShot1:           1054,
			DailyShot2:           107,
			DayDiff:              -39,
			DayTotal:             1161,
			ReferenceDate:        "2021-04-22T00:00:00",
			TotalDistinctPersons: 51872,
			TotalShot1:           51872,
			TotalShot2:           17531,
			TotalVaccinations:    69403,
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
			name: "should fail if API returns an error",
			client: api.NewClient(api.MockAPIToken,
				api.SetHTTPClient(api.NewMockHTTPClient(api.MockRequest{
					Path:  "mdg_emvolio",
					Query: fmt.Sprintf("date_to=%s", time.Now().Format("2006-01-02")),
					Err:   errors.New("cannot fetch data"),
				})),
			),
			params:      api.NewDefaultGetParams(),
			expectedErr: errors.New("get vaccination data: cannot fetch data"),
		},
		{
			name: "should succeed when API call succeeds",
			client: api.NewClient(api.MockAPIToken,
				api.SetHTTPClient(api.NewMockHTTPClient(api.MockRequest{
					Path:         "mdg_emvolio",
					Query:        fmt.Sprintf("date_to=%s", time.Now().Format("2006-01-02")),
					ResponseBody: api.NewMockBody(sampleData),
				})),
			),
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
