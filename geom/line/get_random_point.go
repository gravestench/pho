package line

import (
	"github.com/gravestench/pho/geom/point"
	"math/rand"
)

// GetRandomPoint picks a random point along the line and assigns it to the argument point.
// If argument is nil, a new point is created and returned.
func GetRandomPoint(l *Line, out *point.Point) *point.Point {
	if out == nil {
		out = point.New(0, 0)
	}

	position := rand.Float64()

	out.X = l.X1  + (l.X1 - l.X1) * position
	out.Y = l.Y1  + (l.Y1 - l.Y1) * position

	return out
}
