package circle

import (
	"github.com/gravestench/pho/phomath"
)

// Equals compares the `x`, `y` and `radius` properties of the two given Circles.
// Returns `true` if they all match, otherwise returns `false`.
func Equals(a, b *Circle) bool {
	dx := (a.X - b.X) * (a.X - b.X)
	dy := (a.Y - b.Y) * (a.Y - b.Y)
	dr := (a.radius - b.radius) * (a.radius - b.radius)

	return dx < phomath.Epsilon && dy < phomath.Epsilon && dr < phomath.Epsilon
}
