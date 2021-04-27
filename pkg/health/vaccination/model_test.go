package vaccination

import (
	"testing"

	"github.com/ppapapetrou76/go-testing/assert"
)

func TestList_FilterByArea(t *testing.T) {
	tests := []struct {
		name     string
		initial  List
		expected List
		filter   string
	}{
		{
			name:     "should return empty list if no items match the given filter",
			initial:  List{{Area: "ΑΤΤΙΚΗΣ"}, {Area: "ΒΟΙΩΤΙΑΣ"}},
			expected: List{},
			filter:   "ΘΕΣΣΑΛΟΝΙΚΗΣ",
		},
		{
			name:     "should return the given list if the area name argument is empty",
			initial:  List{{Area: "ΑΤΤΙΚΗΣ"}, {Area: "ΒΟΙΩΤΙΑΣ"}},
			expected: List{{Area: "ΑΤΤΙΚΗΣ"}, {Area: "ΒΟΙΩΤΙΑΣ"}},
			filter:   "",
		},
		{
			name:     "should return a list with the items that match the given filter",
			initial:  List{{Area: "ΑΤΤΙΚΗΣ"}, {Area: "ΒΟΙΩΤΙΑΣ"}},
			expected: List{{Area: "ΑΤΤΙΚΗΣ"}},
			filter:   "ΑΤΤΙΚΗΣ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.That(t, tt.initial.FilterByArea(tt.filter)).IsEqualTo(tt.expected)
		})
	}
}
