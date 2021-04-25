package internal

import (
	"errors"
	"fmt"
	"time"

	"github.com/ppapapetrou76/go-data-gov-gr-sdk/api"
)

// NewCommonMockClientError returns a mock client that handles a single request that fails.
func NewCommonMockClientError(path, errText string) *api.Client {
	return api.NewClient(api.MockAPIToken,
		api.SetHTTPClient(api.NewMockHTTPClient(api.MockRequest{
			Path:  path,
			Query: fmt.Sprintf("date_to=%s", time.Now().Format("2006-01-02")),
			Err:   errors.New(errText),
		})),
	)
}

// NewCommonMockClientSuccess returns a mock client that handles a single request that succeeds.
func NewCommonMockClientSuccess(path, responseBody string) *api.Client {
	return api.NewClient(api.MockAPIToken,
		api.SetHTTPClient(api.NewMockHTTPClient(api.MockRequest{
			Path:         path,
			Query:        fmt.Sprintf("date_to=%s", time.Now().Format("2006-01-02")),
			ResponseBody: api.NewMockBody(responseBody),
		})),
	)
}
