package api

// EndPoint defines a simple interface to describe all available endpoints on data.gov.gr.
type EndPoint interface {
	GetPath() string
	GetName() string
}
