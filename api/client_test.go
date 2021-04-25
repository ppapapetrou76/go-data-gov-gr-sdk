package api

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
	"testing/iotest"
	"time"

	"github.com/ppapapetrou76/go-testing/assert"
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
		queryParams          []QueryParam
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
			queryParams: []QueryParam{
				NewDateQueryParam(time.Now()).AsQueryParam("some_date"),
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

func TestClient_ProcessGetRequest(t *testing.T) {
	type sampleData struct {
		Data string `json:"data"`
	}

	tests := []struct {
		name        string
		params      *GetParams
		client      *Client
		expectedErr error
		expected    sampleData
	}{
		{
			name: "should fail if API returns an error",
			client: NewClient(MockAPIToken,
				SetHTTPClient(NewMockHTTPClient(MockRequest{
					Path:  "some-path",
					Query: fmt.Sprintf("date_to=%s", time.Now().Format("2006-01-02")),
					Err:   errors.New("cannot fetch data"),
				})),
			),
			params:      NewDefaultGetParams(),
			expectedErr: errors.New("some-name: cannot fetch data"),
		},
		{
			name: "should fail if the API response cannot be un-marshaled",
			client: NewClient(MockAPIToken,
				SetHTTPClient(NewMockHTTPClient(MockRequest{
					Path:  "some-path",
					Query: fmt.Sprintf("date_to=%s", time.Now().Format("2006-01-02")),
				})),
			),
			params:      NewDefaultGetParams(),
			expectedErr: errors.New("some-name: failed to un-marshal response unexpected end of JSON input"),
		},
		{
			name: "should fail if the API returns an unexpected response",
			client: NewClient(MockAPIToken,
				SetHTTPClient(NewMockHTTPClient(MockRequest{
					StatusCode:   401,
					Path:         "some-path",
					Query:        fmt.Sprintf("date_to=%s", time.Now().Format("2006-01-02")),
					ResponseBody: NewMockBody(`{"details":"λανθασμένο token"}`),
				})),
			),
			params:      NewDefaultGetParams(),
			expectedErr: errors.New("some-name: {\"details\":\"λανθασμένο token\"}"),
		},
		{
			name: "should fail if the API returns an unexpected response and body cannot be read",
			client: NewClient(MockAPIToken,
				SetHTTPClient(NewMockHTTPClient(MockRequest{
					StatusCode:   401,
					Path:         "some-path",
					Query:        fmt.Sprintf("date_to=%s", time.Now().Format("2006-01-02")),
					ResponseBody: ioutil.NopCloser(iotest.ErrReader(errors.New("cannot read response data"))),
				})),
			),
			params:      NewDefaultGetParams(),
			expectedErr: errors.New("some-name: cannot read response data"),
		},
		{
			name: "should fail if the API returns the expected response and body cannot be read",
			client: NewClient(MockAPIToken,
				SetHTTPClient(NewMockHTTPClient(MockRequest{
					StatusCode:   200,
					Path:         "some-path",
					Query:        fmt.Sprintf("date_to=%s", time.Now().Format("2006-01-02")),
					ResponseBody: ioutil.NopCloser(iotest.ErrReader(errors.New("cannot read response data"))),
				})),
			),
			params:      NewDefaultGetParams(),
			expectedErr: errors.New("some-name: failed to parse response cannot read response data"),
		},
		{
			name: "should return the expected data",
			client: NewClient(MockAPIToken,
				SetHTTPClient(NewMockHTTPClient(MockRequest{
					Path:         "some-path",
					Query:        fmt.Sprintf("date_to=%s", time.Now().Format("2006-01-02")),
					ResponseBody: NewMockBody(`{"data":"value"}`),
				})),
			),
			params: NewDefaultGetParams(),
			expected: sampleData{
				Data: "value",
			},
		},
		{
			name: "should return the expected with query params",
			client: NewClient(MockAPIToken,
				SetHTTPClient(NewMockHTTPClient(MockRequest{
					Path: "some-path",
					Query: fmt.Sprintf("date_from=%s&date_to=%s",
						time.Now().Add(-7*time.Hour*24).Format("2006-01-02"),
						time.Now().Format("2006-01-02")),
					ResponseBody: NewMockBody(`{"data":"value"}`),
				})),
			),
			params: NewDefaultGetParams(SetDateFrom(time.Now().Add(-7 * time.Hour * 24))),
			expected: sampleData{
				Data: "value",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var actual sampleData
			err := tt.client.ProcessGetRequest(tt.params, sampleEndPoint{}, &actual)
			assert.ThatError(t, err).IsSameAs(tt.expectedErr)
			assert.That(t, actual).IsEqualTo(tt.expected)
		})
	}
}

type sampleEndPoint struct{}

func (v sampleEndPoint) GetPath() string {
	return "some-path"
}

func (v sampleEndPoint) GetName() string {
	return "some-name"
}
