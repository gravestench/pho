package line

import (
	"github.com/gravestench/pho/geom/point"
	"github.com/gravestench/pho/phomath"
	"math"
)

// GetNormal calculates the normal of the given line.
// The normal of a line is a vector that points perpendicular from it.
func GetNormal(l *Line, p, out *point.Point) *point.Point {
	if out == nil {
		out = point.New(0, 0)
	}

	a := Angle(l) - phomath.TAU

	out.X = math.Cos(a)
	out.Y = math.Sin(a)

	return out
}
