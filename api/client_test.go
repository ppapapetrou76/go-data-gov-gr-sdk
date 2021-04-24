package api

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/ppapapetrou76/go-testing/assert"

	"github.com/ppapapetrou76/go-data-gov-gr-sdk/internal/httputil"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name, token string
		expected    *Client
		opts        []Opts
	}{
		{
			name:  "should return the default client",
			token: MockAPIToken,
			expected: &Client{
				baseURL:   "https://data.gov.gr/api/v1/query",
				authToken: MockAPIToken,
				timeout:   time.Second * 30,
				httpClient: &http.Client{
					Timeout: time.Second * 30,
				},
			},
		},
		{
			name: "should return the client with opts",
			opts: []Opts{
				SetTimeout(time.Minute * 2),
				SetBaseURL("http://localhost"),
				SetHTTPClient(&MockHTTPClient{}),
			},
			token: MockAPIToken,
			expected: &Client{
				baseURL:    "http://localhost",
				authToken:  MockAPIToken,
				timeout:    time.Minute * 2,
				httpClient: &MockHTTPClient{},
			},
		},
	}

	for _, tt := range tests {
		assert.That(t, NewClient(tt.token, tt.opts...)).IsEqualTo(tt.expected)
	}
}

func TestClient_DoGet(t *testing.T) {
	tests := []struct {
		name, token          string
		client               *Client
		path                 string
		queryParams          []httputil.QueryParam
		expectedResponseBody string
		expectedErr          error
	}{
		{
			name:        "should fail if request generation fails",
			token:       MockAPIToken,
			client:      NewClient(MockAPIToken, SetBaseURL(" http://foo.com")),
			expectedErr: errors.New("failed to create new get request: parse \" http://foo.com/\": first path segment in URL cannot contain colon"),
		},
		{
			name:  "should fail if request sent fails",
			token: MockAPIToken,
			client: NewClient(MockAPIToken, SetHTTPClient(NewMockHTTPClient(
				MockRequest{
					Err: errors.New("cannot send request"),
				},
			))),
			expectedErr: errors.New("cannot send request"),
		},
		{
			name:  "should succeed if request sent succeeds",
			token: MockAPIToken,
			client: NewClient(MockAPIToken, SetHTTPClient(NewMockHTTPClient(
				MockRequest{
					ResponseBody: NewMockBody("{}"),
				},
			))),
			expectedResponseBody: "{}",
		},
		{
			name:  "should succeed if request sent succeeds with query params",
			token: MockAPIToken,
			client: NewClient(MockAPIToken, SetHTTPClient(NewMockHTTPClient(
				MockRequest{
					Query:        fmt.Sprintf("some_date=%s", time.Now().Format("2006-01-02")),
					ResponseBody: NewMockBody("{}"),
				},
			))),
			expectedResponseBody: "{}",
			queryParams: []httputil.QueryParam{
				httputil.NewDateQueryParam(time.Now()).AsQueryParam("some_date"),
				{},
			},
		},
	}

	for _, tt := range tests {
		resp, err := tt.client.DoGet(context.Background(), tt.path, tt.queryParams...)
		assert.ThatError(t, err).IsSameAs(tt.expectedErr)
		if resp != nil {
			actualBody, _ := io.ReadAll(resp.Body)
			resp.Body.Close()
			assert.That(t, string(actualBody)).IsEqualTo(tt.expectedResponseBody)
		}
	}
}
