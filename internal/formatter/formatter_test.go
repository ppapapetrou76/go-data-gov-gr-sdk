package formatter

import (
	"bytes"
	"testing"

	"github.com/ppapapetrou76/go-testing/assert"
)

func TestNew(t *testing.T) {
	var b bytes.Buffer
	tests := []struct {
		name          string
		expected      Formatter
		formatterName string
	}{
		{
			name:     "should create the default formatter (json)",
			expected: NewJSON(&b),
		},
		{
			name:          "should create the json formatter",
			expected:      NewJSON(&b),
			formatterName: "json",
		},
		{
			name:          "should create the yaml formatter",
			expected:      NewYAML(&b),
			formatterName: "yaml",
		},
		{
			name:          "should create the text formatter",
			expected:      NewText(&b),
			formatterName: "text",
		},
	}

	for _, tt := range tests {
		actual := New(&b, tt.formatterName)
		assert.That(t, actual).IsEqualTo(tt.expected)
	}
}
