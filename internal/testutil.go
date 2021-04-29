package internal

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"time"

	"github.com/urfave/cli/v2"

	"github.com/ppapapetrou76/go-data-gov-gr-sdk/api"
)

// AuthTokenRequiredMsg is the message returned by the CLI when the auth-token flag is not set.
const AuthTokenRequiredMsg = `Required flag "auth-token" not set`

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

// NewFlagSet returns the base flag set to be added to the CLI contexts for tests running.
func NewFlagSet() *flag.FlagSet {
	flagSet := &flag.FlagSet{}
	flagSet.String("auth-token", api.MockAPIToken, "")

	return flagSet
}

// NewCLIContext returns a simple CLI context to run tests against it.
func NewCLIContext(flagSet *flag.FlagSet) *cli.Context {
	cliApp := &cli.App{Writer: bytes.NewBufferString("")}

	return cli.NewContext(cliApp, flagSet, nil)
}
