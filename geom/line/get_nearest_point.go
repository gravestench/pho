package line

import (
	"github.com/gravestench/pho/geom/point"
	"github.com/gravestench/pho/phomath"
)

// GetNearestPoint get the nearest point on a line perpendicular to the given point.
func GetNearestPoint(l *Line, p, out *point.Point) *point.Point {
	if out == nil {
		out = point.New(0, 0)
	}

	x1, y1, x2, y2 := l.X1, l.Y1, l.X2, l.Y2

	l2 := ((x2 - x1) * (x2 - x1)) + ((y2 - y1) * (y2 - y1))

	if l2 < phomath.Epsilon {
		return out
	}

	r := (((p.X - x1) * (x2 - x1)) + ((p.Y - y1) * (y2 - y1))) / l2

	out.X = x1 + (r * (x2 - x1))
	out.Y = y1 + (r * (y2 - y1))

	return out
}
