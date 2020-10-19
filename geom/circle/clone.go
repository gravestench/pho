package circle

// Clone returns a clone of this circle
func Clone(c *Circle) *Circle {
	return New(c.X, c.Y, c.radius)
}
