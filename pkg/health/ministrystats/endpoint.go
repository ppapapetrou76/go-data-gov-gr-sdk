package ministrystats

type endpoint struct {
	path, name string
}

// GetPath returns the endpoint path.
func (v endpoint) GetPath() string {
	return v.path
}

// GetName returns the endpoint name.
func (v endpoint) GetName() string {
	return v.name
}
