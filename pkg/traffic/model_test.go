package traffic

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
			initial:  List{{RoadName: "ΚΑΤΕΧΑΚΗ"}, {RoadName: "ΚΗΦΙΣΙΑΣ"}},
			expected: List{},
			filter:   "ΣΥΓΓΡΟΥ",
		},
		{
			name:     "should return a list with the items that match the given filter",
			initial:  List{{RoadName: "ΚΑΤΕΧΑΚΗ"}, {RoadName: "ΚΗΦΙΣΙΑΣ"}},
			expected: List{{RoadName: "ΚΑΤΕΧΑΚΗ"}},
			filter:   "ΚΑΤΕΧΑΚΗ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.That(t, tt.initial.FilterByRoadName(tt.filter)).IsEqualTo(tt.expected)
		})
	}
}
