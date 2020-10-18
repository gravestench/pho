package circle

// Contains checks to see if the Circle contains all four points of the given Rectangle.
func Contains(c *Circle, x, y float64) bool {
	//  Check if x/y are within the bounds first
	if c.radius > 0 && x >= c.Left() && x <= c.Right() && y >= c.Top() && y <= c.Bottom() {
		dx := (c.X - x) * (c.X - x)
		dy := (c.Y - y) * (c.Y - y)

		return (dx + dy) <= (c.radius * c.radius)
	}

	return false
}
