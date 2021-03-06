package circle

import (
	"github.com/gravestench/pho/geom/point"
	"github.com/gravestench/pho/phomath"
)

const (
	ByStepRate = -1
)

// GetPoints Returns an array of Point objects containing the coordinates of the points around the
// circumference of the Circle, based on the given quantity or stepRate values.
func GetPoints(c *Circle, quantity int, stepRate float64, points []*point.Point) []*point.Point {
	if quantity == ByStepRate {
		quantity = int(Circumference(c) / stepRate)
	}

	if points == nil {
		points = make([]*point.Point, 0)
	}

	for idx := 0; idx < quantity; idx++ {
		percent := float64(idx) / float64(quantity)
		angle := phomath.FromPercent(percent, 0, phomath.PI2)

		points = append(points, CircumferencePoint(c, angle, nil))
	}

	return points
}
