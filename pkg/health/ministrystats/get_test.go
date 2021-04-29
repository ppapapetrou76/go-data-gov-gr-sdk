package ministrystats

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
		params      GetParams
		client      *api.Client
		expected    List
		expectedErr error
	}{
		{
			name:   "should fail if stats category is not valid",
			client: internal.NewCommonMockClientError("minhealth_pharmacists", "cannot fetch data"),
			params: GetParams{
				GetParams: api.NewDefaultGetParams(),
				Category:  Category("invalid"),
			},
			expectedErr: errors.New("ministry stats get validation: category invalid is not recognized. valid categories are: [pharmacists pharmacies doctors dentists]"),
		},
		{
			name:   "should fail if API returns an error",
			client: internal.NewCommonMockClientError("minhealth_pharmacists", "cannot fetch data"),
			params: GetParams{
				GetParams: api.NewDefaultGetParams(),
				Category:  Category("pharmacists"),
			},
			expectedErr: errors.New("get pharmacists: cannot fetch data"),
		},
		{
			name:   "should succeed when API call succeeds",
			client: internal.NewCommonMockClientSuccess("minhealth_pharmacists", sampleData),
			params: GetParams{
				GetParams: api.NewDefaultGetParams(),
				Category:  Category("pharmacists"),
			},
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
