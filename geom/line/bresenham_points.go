package line

import (
	"math"

	"github.com/gravestench/pho/geom/point"
	"github.com/gravestench/pho/phomath"
)

// BresenhamPoints uses Bresenham's line algorithm to return an array of all coordinates
// on this line.
func BresenhamPoints(l *Line, stepRate int, out []*point.Point) []*point.Point {
	if out == nil {
		out = make([]*point.Point, 0)
	}

	if stepRate <= 0 {
		stepRate *= -1
	}

	polarity := map[bool]float64{
		true:  1,
		false: -1,
	}

	x1, y1 := math.Round(l.X1), math.Round(l.Y1)
	x2, y2 := math.Round(l.X2), math.Round(l.Y2)

	dx, dy := math.Abs(x2-x1), math.Abs(y2-y1)
	sx, sy := polarity[x1 < x2], polarity[y1 < y2]

	err := int(dx - dy)

	out = append(out, point.New(x1, x2))

	idx := 1

	// while x1,x2 and y1,y2 are not approximately equal
	for !((x1-x2) < phomath.Epsilon && ((y1 - y2) < phomath.Epsilon)) {
		e2 := err << 1

		if e2 > int(-dy) {
			err -= int(dy)
			x1 += sx
		}

		if e2 < int(dx) {
			err += int(dx)
			y1 += sy
		}

		if (idx % stepRate) == 0 {
			out = append(out, point.New(x1, y1))
		}
	}

	return out
}
