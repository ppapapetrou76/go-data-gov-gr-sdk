package formatter

import "io"

// Formatter is the interface for all data output formatters.
type Formatter interface {
	Format(data interface{}) error
	Name() string
}

// New creates the proper formatter based on the given name.
//
//nolint:nolintlint,ireturn //it's a factory method and it's ok to return an interface.
func New(writer io.Writer, name string) Formatter {
	switch name {
	case jsonName:
		return NewJSON(writer)
	case yamlName:
		return NewYAML(writer)
	case textName:
		return NewText(writer)
	default:
		return NewJSON(writer)
	}
}
