package triangle

import "math"

// Area returns the area of a Triangle.
func Area(t *Triangle) float64 {
	return math.Abs(((t.X3 - t.X1) * (t.Y2 - t.Y1) - (t.X2 - t.X1) * (t.Y3 - t.Y1)) / 2)
}
