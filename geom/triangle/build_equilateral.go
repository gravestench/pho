package triangle

import "math"

// BuildEquilateral builds an equilateral triangle. In the equilateral triangle, all the sides are
// the same length (congruent) and all the angles are the same size (congruent). The x/y specifies
// the top-middle of the triangle (x1/y1) and length is the length of each side.
func BuildEquilateral(x, y, length float64) *Triangle {
	height := length * (math.Sqrt(3) / 2)
	x1, y1 := x, y
	x2, y2 := x + (length / 2), y + height
	x3, y3 := x - (length / 2), y + height

	return New(x1, y1, x2, y2, x3, y3)
}
