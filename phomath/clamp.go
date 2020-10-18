package phomath

import "math"

// Clamp will force a value within the boundaries by clamping it to the range `min`, `max`.
func Clamp(value, min, max float64) float64 {
	return math.Max(min, math.Min(max, value))
}
