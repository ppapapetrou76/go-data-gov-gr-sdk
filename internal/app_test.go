package internal

import (
	"net/http"
	"testing"

	"github.com/ppapapetrou76/go-testing/assert"

	"github.com/ppapapetrou76/go-data-gov-gr-sdk/api"
)

func TestInstance(t *testing.T) {
	mockHTTPClient := api.NewMockHTTPClient()
	instance := Instance(mockHTTPClient)

	assert.That(t, instance.HTTPClient).IsEqualTo(mockHTTPClient)
	assert.That(t, Get().HTTPClient).IsEqualTo(mockHTTPClient)

	// try to run the instance again and check that it's not re-initialized ( singleton pattern )
	instance = Instance(&http.Client{})
	assert.That(t, instance.HTTPClient).IsEqualTo(mockHTTPClient)
	assert.That(t, Get().HTTPClient).IsEqualTo(mockHTTPClient)
}
