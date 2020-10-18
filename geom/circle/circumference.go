package circle

import (
	"../../phomath"
)

// Circumference returns the circumference of the given Circle.
func Circumference(c *Circle) float64 {
	return 2 * (phomath.PI * c.radius)
}
