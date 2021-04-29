package ministrystats

import (
	"fmt"

	"github.com/ppapapetrou76/go-data-gov-gr-sdk/api"
)

// GetParams is used by the Get function.
type GetParams struct {
	*api.GetParams
	Category
}

func (params GetParams) validate() error {
	if !params.Category.isValid() {
		return fmt.Errorf("ministry stats get validation: category %s is not recognized. valid categories are: %+v",
			params.Category, validCategories())
	}

	return nil
}

// Get returns a list of Attica traffic data found based on the given parameters.
func Get(client *api.Client, params GetParams) (List, error) {
	if err := params.validate(); err != nil {
		return nil, err
	}
	var results List
	if err := client.ProcessGetRequest(params.GetParams, endpoint{
		path: fmt.Sprintf("%s_%s", endPointPrefix, params.Category),
		name: string(params.Category),
	}, &results); err != nil {
		return nil, fmt.Errorf("%s %w", "get", err)
	}

	return results, nil
}
