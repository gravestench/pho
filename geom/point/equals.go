package point

import (
	"github.com/gravestench/pho/phomath"
	"math"
)

// Equals compares two `Point` objects to see if they are equal.
func Equals(a, b *Point) bool {
	dx, dy := math.Abs(a.X -b.X), math.Abs(a.Y -b.Y)

	return dx < phomath.Epsilon &&  dy < phomath.Epsilon
}
