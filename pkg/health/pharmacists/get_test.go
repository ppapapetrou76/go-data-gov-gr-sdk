package pharmacist

import (
	"errors"
	"testing"

	"github.com/ppapapetrou76/go-testing/assert"

	"github.com/ppapapetrou76/go-data-gov-gr-sdk/api"
	"github.com/ppapapetrou76/go-data-gov-gr-sdk/internal"
)

const sampleData = `[
 {
    "year": 2018,
    "quarter": "Q3",
    "active": 10616,
    "entrants": 30,
    "exits": 53
  },
  {
    "year": 2018,
    "quarter": "Q4",
    "active": 10587,
    "entrants": 40,
    "exits": 69
  }
]`

func expectedResult() List {
	return []Data{
		{
			Year:     2018,
			Quarter:  "Q3",
			Active:   10616,
			Entrants: 30,
			Exits:    53,
		},
		{
			Year:     2018,
			Quarter:  "Q4",
			Active:   10587,
			Entrants: 40,
			Exits:    69,
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
			client:      internal.NewCommonMockClientError("minhealth_pharmacists", "cannot fetch data"),
			params:      api.NewDefaultGetParams(),
			expectedErr: errors.New("get pharmacists data: cannot fetch data"),
		},
		{
			name:     "should succeed when API call succeeds",
			client:   internal.NewCommonMockClientSuccess("minhealth_pharmacists", sampleData),
			params:   api.NewDefaultGetParams(),
			expected: expectedResult(),
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
