package line

import "github.com/gravestench/pho/geom/point"

// GetMidPoint get the midpoint of the given line. Assigns to the `out` point,
// or creates one if nil.
func GetMidPoint(l *Line, out *point.Point) *point.Point {
	if out == nil {
		out = point.New(0, 0)
	}

	out.X = (l.X1 + l.X2) / 2
	out.Y = (l.Y1 + l.Y2) / 2

	return out
}
