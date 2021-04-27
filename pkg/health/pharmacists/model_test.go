package pharmacist

import (
	"testing"

	"github.com/ppapapetrou76/go-testing/assert"
)

func TestList_FilterByArea(t *testing.T) {
	tests := []struct {
		name             string
		initial          List
		expected         List
		fromYear, toYear int
	}{
		{
			name:     "should return empty list if no items match the given filter",
			initial:  List{{Year: 2020}, {Year: 2021}, {Year: 2022}},
			expected: List{},
			fromYear: 1990,
			toYear:   2000,
		},
		{
			name:     "should return a list with the items that match the given filter",
			initial:  List{{Year: 2020}, {Year: 2021}, {Year: 2022}},
			expected: List{{Year: 2021}},
			fromYear: 2021,
			toYear:   2021,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.That(t, tt.initial.FilterByYearRange(tt.fromYear, tt.toYear)).IsEqualTo(tt.expected)
		})
	}
}
