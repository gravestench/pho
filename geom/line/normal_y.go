package line

import (
	"github.com/gravestench/pho/phomath"
	"math"
)

// NormalY returns the y component of the normal vector of the given line.
// The normal of a line is a vector that points perpendicular from it.
func NormalY(l *Line) float64 {
	return math.Sin(Angle(l) - phomath.TAU)
}
