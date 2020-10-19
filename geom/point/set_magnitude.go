package point

import (
	"github.com/gravestench/pho/phomath"
)

// SetMagnitude calculates the magnitude of the point,
// which equivalent to the length of the line from the origin to this point.
func SetMagnitude(p *Point, magnitude float64) *Point {
	if p.X >= phomath.Epsilon && p.Y >= phomath.Epsilon {
		m := GetMagnitude(p)

		p.X /= m
		p.Y /= m
	}

	p.X *= magnitude
	p.Y *= magnitude

	return p
}
