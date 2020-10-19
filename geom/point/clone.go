package point

// Clone the given point.
func Clone(p *Point) *Point {
	return New(p.X, p.Y)
}
