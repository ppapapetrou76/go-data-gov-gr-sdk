package api

import "time"

// GetParams is used by the Get function.
type GetParams struct {
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
func NewDefaultGetParams(opts ...GetParamsOpts) *GetParams {
	now := time.Now()
	params := GetParams{
		dateTo: now,
	}

	for _, o := range opts {
		o(&params)
	}

	return &params
}
