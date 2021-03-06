package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/ppapapetrou76/go-testing/assert"
	"github.com/urfave/cli/v2"

	"github.com/ppapapetrou76/go-data-gov-gr-sdk/api"
	"github.com/ppapapetrou76/go-data-gov-gr-sdk/internal"
)

func TestCommands(t *testing.T) {
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
		name           string
		args           []string
		expectedErr    error
		expectedOutput string
		httpClient     api.HTTPClient
	}{
		{
			name:        "should error if auth-token is not provided",
			args:        []string{"", "health", "ministry-stats"},
			expectedErr: errors.New(internal.AuthTokenRequiredMsg),
		},
		{
			name: "should error if getting ministry-stats data fails",
			httpClient: api.NewMockHTTPClient(api.MockRequest{
				Path:  "minhealth_pharmacists",
				Query: fmt.Sprintf("date_to=%s", time.Now().Format("2006-01-02")),
				Err:   errors.New("cannot send request"),
			}),
			args:        []string{"", "health", "ministry-stats", "--auth-token", api.MockAPIToken, "-c", "pharmacists"},
			expectedErr: errors.New("ministry statistics:get pharmacists: cannot send request"),
		},
		{
			name: "should display data in the default json format",
			httpClient: api.NewMockHTTPClient(api.MockRequest{
				Path:         "minhealth_pharmacists",
				Query:        fmt.Sprintf("date_to=%s", time.Now().Format("2006-01-02")),
				ResponseBody: api.NewMockBody(sampleData),
			}),
			args:           []string{"", "health", "ministry-stats", "--auth-token", api.MockAPIToken, "-c", "pharmacists"},
			expectedOutput: sampleData,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			internal.Get().HTTPClient = tt.httpClient
			writer := bytes.NewBufferString("")
			cliApp := &cli.App{
				Writer: writer,
			}
			cliApp.Commands = Commands()
			err := cliApp.Run(tt.args)

			assert.ThatError(t, err).IsSameAs(tt.expectedErr)
			if tt.expectedOutput != "" && tt.expectedErr == nil {
				assert.That(t, writer.String()).IsEqualTo(tt.expectedOutput)
			}
		})
	}
}
