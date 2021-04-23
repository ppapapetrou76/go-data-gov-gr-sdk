package api

import (
	"time"
)

// Opts is function type that implements the Opt pattern to dynamically set Client attributes.
type Opts func(*Client)

// SetTimeout returns an Opts function that sets the http client request timeout.
func SetTimeout(timeout time.Duration) Opts {
	return func(client *Client) {
		client.timeout = timeout
	}
}

// SetBaseURL returns an Opts function that sets the http client base URL.
func SetBaseURL(baseURL string) Opts {
	return func(client *Client) {
		client.baseURL = baseURL
	}
}

// SetHTTPClient returns an Opts function that sets the http client client.
func SetHTTPClient(httpClient HTTPClient) Opts {
	return func(client *Client) {
		client.httpClient = httpClient
	}
}
