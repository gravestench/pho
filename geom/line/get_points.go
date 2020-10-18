package line

import "github.com/gravestench/pho/geom/point"

const (
	AutoQuantity = -1
	AutoStepRate = -1
)

// Get a number of points along a line's length.
// Provide a `quantity` to get an exact number of points along the line.
func GetPoints(l *Line, quantity int, stepRate float64, out []*point.Point) []*point.Point {
	if quantity <= 0 && stepRate > 0 {
		quantity = int(Length(l) / stepRate)
	}

	if out == nil {
		out = make([]*point.Point, 0)
	}

	for idx := 0; idx < quantity; idx++ {
		position := float64(idx) / float64(quantity)

		x := l.X1 + (l.X2-l.X1)*position
		y := l.Y1 + (l.Y2-l.Y1)*position

		out = append(out, point.New(x, y))
	}

	return out
}
