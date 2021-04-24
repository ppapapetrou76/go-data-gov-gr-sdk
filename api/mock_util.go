package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"reflect"

	"github.com/stretchr/testify/mock"
)

// MockAPIToken holds the value of a fake API token.
const MockAPIToken = "mock-token"

var defaultHeaders = http.Header{
	"Content-Type":  {defaultContentType},
	"Accept":        {defaultContentType},
	"Authorization": {fmt.Sprintf("Token %s", MockAPIToken)},
}

// MockRequest is used to mock an API call.
type MockRequest struct {
	Path, Query  string
	StatusCode   int
	Err          error
	ResponseBody io.ReadCloser
}

// NewMockBody returns a mocked response body with the given response data.
func NewMockBody(response string) io.ReadCloser {
	return io.NopCloser(bytes.NewBuffer([]byte(response)))
}

// NewMockHTTPClient is a utility function to return a mocked http client with the given list of
// requests.
func NewMockHTTPClient(requests ...MockRequest) *MockHTTPClient {
	mockedClient := &MockHTTPClient{}
	for _, r := range requests {
		if r.StatusCode == 0 {
			r.StatusCode = http.StatusOK
		}
		if r.ResponseBody == nil {
			r.ResponseBody = io.NopCloser(bytes.NewBuffer([]byte("")))
		}

		path := fmt.Sprintf("%s/%s", defaultURL, r.Path)
		if r.Query != "" {
			path = fmt.Sprintf("%s?%s", path, r.Query)
		}
		mockedClient.On("Do", mock.MatchedBy(
			httpRequestMatcher(http.MethodGet, path, defaultHeaders)),
		).Return(&http.Response{StatusCode: r.StatusCode, Body: r.ResponseBody}, r.Err)
	}
	return mockedClient
}

func httpRequestMatcher(method, path string, headers http.Header) func(r *http.Request) bool {
	return func(r *http.Request) bool {
		return reflect.DeepEqual(headers, r.Header) &&
			r.Method == method &&
			r.URL.String() == path
	}
}
