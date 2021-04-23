package httputil

import "time"

const dateFormat = "2006-01-02"

// QueryParam defines an API query name.
type QueryParam struct {
	name, value string
}

// GetParameter returns the query param name.
func (q QueryParam) GetParameter() string {
	return q.name
}

// GetValue returns the query param value.
func (q QueryParam) GetValue() string {
	return q.value
}

// DateQueryValue is used to hold a time.Time optional value.
type DateQueryValue struct {
	value time.Time
}

// NewDateQueryParam returns a DateQueryValue object initialised with the given value.
func NewDateQueryParam(value time.Time) DateQueryValue {
	return DateQueryValue{
		value: value,
	}
}

// AsQueryParam returns a QueryParam object with the formatted date value. If the value is nil then it returns
// an empty QueryParam object.
func (p DateQueryValue) AsQueryParam(parameter string) QueryParam {
	if p.value.Nanosecond() == 0 {
		return QueryParam{}
	}
	return QueryParam{
		name:  parameter,
		value: p.value.Format(dateFormat),
	}
}
