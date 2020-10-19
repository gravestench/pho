package point

// Invert swaps the X and the Y coordinate of a point.
func Invert(p *Point) *Point {
	return p.SetTo(p.Y, p.X)
}
