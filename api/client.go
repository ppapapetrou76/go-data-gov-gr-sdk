package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	// DefaultTimeout is the default API request timeout.
	DefaultTimeout = time.Minute

	defaultURL         = `https://data.gov.gr/api/v1/query`
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

// NewClient returns a new API client initialised with the given host URL and a list of optional Opts functions.
func NewClient(authToken string, opts ...Opts) *Client {
	client := &Client{
		authToken: authToken,
		timeout:   DefaultTimeout,
		baseURL:   defaultURL,
		httpClient: &http.Client{
			Timeout: DefaultTimeout,
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
func (c *Client) DoGet(ctx context.Context, path string, queryParams ...QueryParam) (*http.Response, error) {
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

// ProcessGetRequest processes a full GET request.
// It prepares the query params, runs the actual API Get request, checks for the expected status code and unmarshal the
// API response to the given structure.
// If any of the above fails then it returns a wrapped error.
func (c *Client) ProcessGetRequest(params *GetParams, endPoint EndPoint, results interface{}) error {
	queryParams := []QueryParam{
		NewDateQueryParam(params.dateFrom).AsQueryParam("date_from"),
		NewDateQueryParam(params.dateTo).AsQueryParam("date_to"),
	}

	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	resp, err := c.DoGet(ctx, endPoint.GetPath(), queryParams...)
	if err != nil {
		return fmt.Errorf("%s: %w", endPoint.GetName(), err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("%s: %w", endPoint.GetName(), err)
		}

		return fmt.Errorf("%s: %s", endPoint.GetName(), string(b))
	}

	if err = parseResponse(resp, &results); err != nil {
		return fmt.Errorf("%s: %w", endPoint.GetName(), err)
	}

	return nil
}

// parseResponse is a utility function to parse a response and unmarshal the json body
// It's not bound to any specific model although it's only used for the Account structure in the scope of this library.
func parseResponse(resp *http.Response, model interface{}) error {
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to parse response %w", err)
	}
	err = json.Unmarshal(b, model)
	if err != nil {
		return fmt.Errorf("failed to un-marshal response %w", err)
	}

	return nil
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
