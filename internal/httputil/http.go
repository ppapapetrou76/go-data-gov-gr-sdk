package httputil

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ParseResponse is a utility function to parse a response and unmarshal the json body
// It's not bound to any specific model although it's only used for the Account structure in the scope of this library.
func ParseResponse(resp *http.Response, model interface{}) error {
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
