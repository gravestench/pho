package phomath

import "math"

// Wrap the given `value` between `min` and `max.
func Wrap(value, min, max float64) float64 {
	r := max - min

	return min + math.Mod(math.Mod(value - min, r) + r, r)
}
