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
	Path, Query, ResponseBody string
	Err                       error
}

// NewMockHTTPClient is a utility function to return a mocked http client with the given list of
// requests.
func NewMockHTTPClient(requests ...MockRequest) *MockHTTPClient {
	mockedClient := &MockHTTPClient{}
	for _, r := range requests {
		path := fmt.Sprintf("%s/%s", defaultURL, r.Path)
		if r.Query != "" {
			path = fmt.Sprintf("%s?%s", path, r.Query)
		}
		mockedClient.On("Do", mock.MatchedBy(
			httpRequestMatcher(http.MethodGet, path, defaultHeaders)),
		).Return(&http.Response{Body: io.NopCloser(bytes.NewBuffer([]byte(r.ResponseBody)))}, r.Err)
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
