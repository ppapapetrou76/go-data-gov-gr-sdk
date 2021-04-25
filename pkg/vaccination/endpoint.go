package vaccination

// Endpoint is vaccination API endpoint.
type Endpoint struct{}

const (
	path = "mdg_emvolio"
	name = "vaccination data"
)

// GetPath returns the endpoint path.
func (v Endpoint) GetPath() string {
	return path
}

// GetName returns the endpoint name.
func (v Endpoint) GetName() string {
	return name
}
