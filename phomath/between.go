package phomath

import (
	"math"
	"math/rand"
)

// Between computes a random integer between the `min` and `max` values, inclusive.
func Between(min, max int) int {
	return int(math.Floor(rand.Float64() *  float64(max - min + 1) + float64(min)))
}
