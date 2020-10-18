package circle

import "math"

// Area returns the circle area
func Area(c *Circle) float64 {
	return math.Pi * c.radius * c.radius
}
