package line

import "github.com/gravestench/pho/geom/point"

// GetPoint on a line that's a given percentage along its length.
func GetPoint(l *Line, position float64, out *point.Point) *point.Point {
	if out == nil {
		out = point.New(0, 0)
	}

	out.X = l.X1 + (l.X2-l.X1)*position
	out.Y = l.Y1 + (l.Y2-l.Y1)*position

	return out
}
