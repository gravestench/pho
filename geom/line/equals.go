package line

import (
	"github.com/gravestench/pho/phomath"
	"math"
)

// Equals checks if a line is approximately equal to another
func Equals(a, b *Line) bool {
	dx1, dy1 := math.Abs(a.X1 - b.X1), math.Abs(a.Y1 - b.Y1)
	dx2, dy2 := math.Abs(a.X2 - b.X2), math.Abs(a.Y2 - b.Y2)

	return dx1 < phomath.Epsilon &&
		dy1 < phomath.Epsilon &&
		dx2 < phomath.Epsilon &&
		dy2 < phomath.Epsilon
}
