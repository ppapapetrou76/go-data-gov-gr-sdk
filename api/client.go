package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ppapapetrou76/go-data-gov-gr-sdk/internal/httputil"
)

const (
	defaultURL         = `https://data.gov.gr/api/v1/query`
	defaultTimeout     = time.Second * 30
	defaultContentType = "application/json"
)

// HTTPClient interface.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client is the structure that initializes the account API Client.
type Client struct {
	baseURL, authToken string
	timeout            time.Duration
	httpClient         HTTPClient
}

// GetTimeout returns the client's timeout value.
func (c Client) GetTimeout() time.Duration {
	return c.timeout
}

// NewClient returns a new API client initialised with the given host URL and a list of optional Opts functions.
func NewClient(authToken string, opts ...Opts) *Client {
	client := &Client{
		authToken: authToken,
		timeout:   defaultTimeout,
		baseURL:   defaultURL,
		httpClient: &http.Client{
			Timeout: defaultTimeout,
		},
	}
	for _, o := range opts {
		o(client)
	}

	return client
}

// DoGet performs an http GET request and returns the outcome of it ( response and an optional error )
// The `path` argument is added to the client base URL and the optional query parameters are encoded and added to the
// final URL too.
func (c *Client) DoGet(ctx context.Context, path string, queryParams ...httputil.QueryParam) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", c.baseURL, path), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create new get request: %w", err)
	}

	q := req.URL.Query()
	for _, p := range queryParams {
		if p.GetParameter() == "" {
			continue
		}
		q.Add(p.GetParameter(), p.GetValue())
	}
	req.URL.RawQuery = q.Encode()
	return c.sendRequest(req.WithContext(ctx))
}

// sendRequest sends an http request using the API client.
// It sets the default headers including the auth token, timeout and executes the request.
// In this example it can be unexported but if we want to add more APIs (in other packages)
// probably this method should be exported and the whole file should be moved to another (common) package.
func (c *Client) sendRequest(req *http.Request) (*http.Response, error) {
	req.Header.Set("Content-Type", defaultContentType)
	req.Header.Set("Accept", defaultContentType)
	req.Header.Set("Authorization", fmt.Sprintf("Token %s", c.authToken))
	//nolint:wrapcheck // it's an unexported method. Wrapping is happening by the client.
	return c.httpClient.Do(req)
}
