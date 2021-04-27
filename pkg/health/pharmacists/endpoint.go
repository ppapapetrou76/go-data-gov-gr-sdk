package pharmacist

// Endpoint is traffic API endpoint.
type Endpoint struct{}

const (
	path = "minhealth_pharmacists"
	name = "pharmacists data"
)

// GetPath returns the endpoint path.
func (v Endpoint) GetPath() string {
	return path
}

// GetName returns the endpoint name.
func (v Endpoint) GetName() string {
	return name
}
