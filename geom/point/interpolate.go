package point

// Interpolate returns the linear interpolation point between the two given points, based on `t`.
func Interpolate(a, b *Point, t float64, out *Point) *Point {
	if out == nil {
		out = New(0, 0)
	}

	out.X = a.X + ((b.X - a.X) * t)
	out.Y = a.Y + ((b.Y - a.Y) * t)

	return out
}
