package point

import "github.com/gravestench/pho/phomath"

// Project calculates the vector projection of `pointA` onto the nonzero `pointB`. This is the
// orthogonal projection of `pointA` onto a straight line parallel to `pointB`.
func Project(a, b, out *Point) *Point {
	if out == nil {
		out = New(0, 0)
	}

	dot := (a.X * b.X) + (a.Y * b.Y)
	amt := dot / GetMagnitudeSquared(b)

	if amt >= phomath.Epsilon {
		out.X = amt * b.X
		out.Y = amt * b.Y
	}

	return out
}
