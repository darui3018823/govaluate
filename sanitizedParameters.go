package govaluate

import "github.com/darui3018823/govaluate/internal/value"

// sanitizedParameters is a wrapper for Parameters that does sanitization as
// parameters are accessed.
type sanitizedParameters struct {
	orig Parameters
}

func (p sanitizedParameters) Get(key string) (interface{}, error) {
	value, err := p.orig.Get(key)
	if err != nil {
		return nil, err
	}

	return normalizeNumber(value), nil
}

func normalizeNumber(input interface{}) interface{} {
	return value.NormalizeNumber(input)
}
