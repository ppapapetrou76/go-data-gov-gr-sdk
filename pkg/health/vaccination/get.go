package vaccination

import (
	"fmt"

	"github.com/ppapapetrou76/go-data-gov-gr-sdk/api"
)

// Get returns a list of vaccinations found based on the given parameters.
func Get(client *api.Client, params *api.GetParams) (List, error) {
	var results List
	if err := client.ProcessGetRequest(params, Endpoint{}, &results); err != nil {
		return nil, fmt.Errorf("%s %w", "get", err)
	}

	return results, nil
}
