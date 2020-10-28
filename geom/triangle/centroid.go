package triangle

import "github.com/gravestench/pho/geom/point"

// Calculates the position of a Triangle's centroid, which is also its center of
// mass ( center of gravity). The centroid is the point in a Triangle at which its
// three medians ( the lines drawn from the vertices to the bisectors of the opposite sides) meet.
// It divides each one in a 2:1 ratio.
func Centroid(t *Triangle, out *point.Point) *point.Point {
	if out == nil {
		out = point.New(0, 0)
	}

	out.X = (t.X1 + t.X2 + t.X3) / 3
	out.Y = (t.Y1 + t.Y2 + t.Y3) / 3

	return out
}
