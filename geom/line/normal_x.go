package line

import (
	"github.com/gravestench/pho/phomath"
	"math"
)

// NormalX returns the x component of the normal vector of the given line.
// The normal of a line is a vector that points perpendicular from it.
func NormalX(l *Line) float64 {
	return math.Cos(Angle(l) - phomath.TAU)
}
