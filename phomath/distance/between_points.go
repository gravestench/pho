package distance

import (
	"github.com/gravestench/pho/geom/point"
	"math"
)

// DistanceBetweenPoints calculates the distance between two points.
func DistanceBetweenPoints(a, b *point.Point) float64 {
	dx, dy := a.X - b.X, a.Y - b.Y

	return math.Sqrt(dx * dx + dy  * dy)
}
