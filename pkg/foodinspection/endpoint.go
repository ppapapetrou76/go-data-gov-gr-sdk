package foodinspection

// Endpoint is the food inspections API endpoint.
type Endpoint struct{}

const (
	path = "efet_inspections"
	name = "food inspections data"
)

// GetPath returns the endpoint path.
func (v Endpoint) GetPath() string {
	return path
}

// GetName returns the endpoint name.
func (v Endpoint) GetName() string {
	return name
}
