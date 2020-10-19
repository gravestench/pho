package line

import (
	"github.com/gravestench/pho/geom/point"
	"github.com/gravestench/pho/phomath"
	"math"
)

// GetShortestDistance get the shortest distance from a Line to the given Point.
func GetShortestDistance(l *Line, p *point.Point) float64 {
	x1, y1, x2, y2 := l.X1, l.Y1, l.X2, l.Y2

	l2 := ((x2 - x1) * (x2 - x1)) + ((y2 - y1) * (y2 - y1))

	if l2 < phomath.Epsilon {
		return l2
	}

	var s = (((y1 - p.Y) * (x2 - x1)) - ((x1 - p.X) * (y2 - y1))) / l2

	return math.Abs(s) * math.Sqrt(l2)
}
