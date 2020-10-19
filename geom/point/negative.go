package point

// Negative flips the sign of the X and the Y coordinate of a point.
func Negative(p *Point) *Point {
	return p.SetTo(-p.X, -p.Y)
}
