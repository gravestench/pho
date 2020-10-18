package circle

// Offset the circle position by the given x, y
func Offset(c *Circle, x, y float64) *Circle {
	c.X += x
	c.Y += y

	return c
}
