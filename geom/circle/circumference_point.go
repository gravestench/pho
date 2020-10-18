package circle

import (
	"math"

	"github.com/gravestench/pho/geom/point"
)

// CircumferencePoint returns a Point containing the coordinates of a point on the
// circumference of the Circle based on the given angle.
func CircumferencePoint(c *Circle, angle float64, p *point.Point) *point.Point {
	if p == nil {
		p = point.New(0, 0)
	}

	p.X = c.X + (c.radius * math.Cos(angle))
	p.Y = c.Y + (c.radius * math.Sin(angle))

	return p
}
