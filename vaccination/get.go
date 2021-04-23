package vaccination

import (
	"context"
	"fmt"
	"time"

	"github.com/ppapapetrou76/go-data-gov-gr-sdk/api"
	"github.com/ppapapetrou76/go-data-gov-gr-sdk/internal/httputil"
)

const (
	path   = "mdg_emvolio"
	prefix = "get vaccination data"
)

// GetParams is used by the Get function.
type GetParams struct {
	client           *api.Client
	dateFrom, dateTo time.Time
}

// GetParamsOpts is function type that implements the Opt pattern to dynamically set GetParams attributes.
type GetParamsOpts func(*GetParams)

// SetDateFrom dynamically sets the `dateFrom` field of GetParams.
func SetDateFrom(dateFrom time.Time) GetParamsOpts {
	return func(params *GetParams) {
		params.dateFrom = dateFrom
	}
}

// SetDateTo dynamically sets the `dateTO` field of GetParams.
func SetDateTo(dateTo time.Time) GetParamsOpts {
	return func(params *GetParams) {
		params.dateTo = dateTo
	}
}

// NewDefaultGetParams returns the address value of a new GetParams initialised with the given api.Client and other
// default values that make the structure usable.
func NewDefaultGetParams(client *api.Client, opts ...GetParamsOpts) *GetParams {
	now := time.Now()
	params := GetParams{
		client: client,
		dateTo: now,
	}

	for _, o := range opts {
		o(&params)
	}
	return &params
}

// Get returns a list of vaccinations found based on the given parameters.
func Get(params *GetParams) (List, error) {
	queryParams := []httputil.QueryParam{
		httputil.NewDateQueryParam(params.dateFrom).AsQueryParam("date_from"),
		httputil.NewDateQueryParam(params.dateTo).AsQueryParam("date_to"),
	}

	ctx, cancel := context.WithTimeout(context.Background(), params.client.GetTimeout())
	defer cancel()

	resp, err := params.client.DoGet(ctx, path, queryParams...)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", prefix, err)
	}

	var results List
	if err = httputil.ParseResponse(resp, &results); err != nil {
		return nil, fmt.Errorf("%s: %w", prefix, err)
	}
	return results, nil
}
