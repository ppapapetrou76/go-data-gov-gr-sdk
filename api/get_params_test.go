package api

import (
	"testing"
	"time"

	"github.com/ppapapetrou76/go-testing/assert"
)

func TestNewDefaultGetParams(t *testing.T) {
	tests := []struct {
		name     string
		expected *GetParams
		opts     []GetParamsOpts
	}{
		{
			name: "should return default get params without opts",
			expected: &GetParams{
				dateTo: time.Now(),
			},
		},
		{
			name: "should return default get params with opts",
			expected: &GetParams{
				dateTo:   time.Now().Add(time.Hour * 24 * 7),
				dateFrom: time.Now().Add(-time.Hour * 24 * 7),
			},
			opts: []GetParamsOpts{
				SetDateFrom(time.Now().Add(-time.Hour * 24 * 7)),
				SetDateTo(time.Now().Add(time.Hour * 24 * 7)),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := NewDefaultGetParams(tt.opts...)
			if len(tt.opts) == 0 {
				assert.ThatTime(t, actual.dateFrom).IsNotDefined()
			} else {
				assert.ThatTime(t, actual.dateFrom).IsAlmostSameAs(tt.expected.dateFrom)
			}
			assert.ThatTime(t, actual.dateTo).IsAlmostSameAs(tt.expected.dateTo)
		})
	}
}
