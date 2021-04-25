package traffic

// Endpoint is traffic API endpoint.
type Endpoint struct{}

const (
	path = "road_traffic_attica"
	name = "traffic data"
)

// GetPath returns the endpoint path.
func (v Endpoint) GetPath() string {
	return path
}

// GetName returns the endpoint name.
func (v Endpoint) GetName() string {
	return name
}
