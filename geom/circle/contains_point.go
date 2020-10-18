package circle

import "github.com/gravestench/pho/geom/point"

// ContainsPoint checks to see if the circle contains the given point
func ContainsPoint(c *Circle, p *point.Point) bool {
	return Contains(c, p.X, p.Y)
}
