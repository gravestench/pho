package circle

import (
	"../../phomath"
	"../point"
)

// GetPoint returns a Point object containing the coordinates of a point on the circumference of
// the Circle based on the given angle normalized to the range 0 to 1. I.e. a value of 0.5 will give
// the point at 180 degrees around the circle.
func GetPoint(c *Circle, position float64, p *point.Point) *point.Point {
	angle := phomath.FromPercent(position, 0, phomath.PI2)

	return CircumferencePoint(c, angle, p)
}
