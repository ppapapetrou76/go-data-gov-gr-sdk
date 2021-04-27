package numbers

import (
	"testing"

	"github.com/ppapapetrou76/go-testing/assert"
)

func TestMin(t *testing.T) {
	tests := []struct {
		name                 string
		one, other, expected int
	}{
		{
			name:     "should return one number",
			one:      1,
			other:    2,
			expected: 1,
		},
		{
			name:     "should return other number",
			one:      10,
			other:    2,
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ThatInt(t, Min(tt.one, tt.other)).IsEqualTo(tt.expected)
		})
	}
}
