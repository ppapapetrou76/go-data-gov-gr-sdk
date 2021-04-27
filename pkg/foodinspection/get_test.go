package foodinspection

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
    "inspections": 8987,
    "violations": 1092,
    "violating_organizations": 646,
    "penalties": 1482227.75
  },
  {
    "year": 2019,
    "inspections": 8586,
    "violations": 1296,
    "violating_organizations": 778,
    "penalties": 1964030.75
  }
]`

func expectedResult() List {
	return []Data{
		{
			Year:                   2018,
			Inspections:            8987,
			Violations:             1092,
			ViolatingOrganizations: 646,
			Penalties:              1482227.75,
		},
		{
			Year:                   2019,
			Inspections:            8586,
			Violations:             1296,
			ViolatingOrganizations: 778,
			Penalties:              1964030.75,
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
			client:      internal.NewCommonMockClientError("efet_inspections", "cannot fetch data"),
			params:      api.NewDefaultGetParams(),
			expectedErr: errors.New("get food inspections data: cannot fetch data"),
		},
		{
			name:     "should succeed when API call succeeds",
			client:   internal.NewCommonMockClientSuccess("efet_inspections", sampleData),
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
