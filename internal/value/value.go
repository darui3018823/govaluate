// Package value contains internal runtime value normalization helpers.
package value

// NormalizeNumber converts numeric parameter values to the float64 shape used
// by govaluate's evaluator.
func NormalizeNumber(input any) any {
	switch value := input.(type) {
	case uint8:
		return float64(value)
	case uint16:
		return float64(value)
	case uint32:
		return float64(value)
	case uint64:
		return float64(value)
	case int8:
		return float64(value)
	case int16:
		return float64(value)
	case int32:
		return float64(value)
	case int64:
		return float64(value)
	case int:
		return float64(value)
	case float32:
		return float64(value)
	}

	return input
}
