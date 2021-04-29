package cmdhealth

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/ppapapetrou76/go-testing/assert"

	"github.com/ppapapetrou76/go-data-gov-gr-sdk/api"
	"github.com/ppapapetrou76/go-data-gov-gr-sdk/internal"
	cmdglobal "github.com/ppapapetrou76/go-data-gov-gr-sdk/internal/cmd/global"
)

func Test_pharmacistCmd(t *testing.T) {
	cmd := pharmacistCmd()

	assert.That(t, cmd.Name).IsEqualTo("pharmacist")
	assert.That(t, cmd.Usage).IsEqualTo("Shows pharmacists statistics")
	assert.ThatSlice(t, cmd.Flags).
		HasSize(4).
		Contains(cmdglobal.CommonFlags()).
		Contains(cmdglobal.YearRangeFlags())
}

func Test_pharmacistCmd_Action(t *testing.T) {
	sampleData := `[
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
]
`
	internal.Instance(nil)
	tests := []struct {
		name        string
		expectedErr error
		httpClient  api.HTTPClient
	}{
		{
			name: "should error if getting pharmacists data fails",
			httpClient: api.NewMockHTTPClient(api.MockRequest{
				Path:  "minhealth_pharmacists",
				Query: fmt.Sprintf("date_to=%s", time.Now().Format("2006-01-02")),
				Err:   errors.New("cannot send request"),
			}),
			expectedErr: errors.New("pharmacists statistics:get pharmacists data: cannot send request"),
		},
		{
			name: "should succeed",
			httpClient: api.NewMockHTTPClient(api.MockRequest{
				Path:         "minhealth_pharmacists",
				Query:        fmt.Sprintf("date_to=%s", time.Now().Format("2006-01-02")),
				ResponseBody: api.NewMockBody(sampleData),
			}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := pharmacistCmd()
			internal.Get().HTTPClient = tt.httpClient
			err := cmd.Action(internal.NewCLIContext())
			assert.ThatError(t, err).IsSameAs(tt.expectedErr)
		})
	}
}
