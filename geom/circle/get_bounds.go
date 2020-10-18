package circle

// GetBounds returns the bounds (a rectangle) of the Circle object.
func GetBounds(c *Circle, out *rectangle.Rectangle) *rectangle.Rectangle {
	if out == nil {
		out = rectangle.New(0, 0, 0, 0)
	}

	out.X, out.Y = c.Left(), c.Top()
	out.Width, out.Height = c.diameter, c.diameter

	return out
}
