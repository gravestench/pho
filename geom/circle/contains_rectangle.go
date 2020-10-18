package circle

import "github.com/gravestench/pho/geom/rectangle"

// ContainsRectangle checks to see if the circle contains the given rectangle
func ContainsRectangle(c *Circle, r *rectangle.Rectangle) bool {
	return Contains(c, r.X, r.Y) &&
		Contains(c, r.Right(), r.Y) &&
		Contains(c, r.X, r.Bottom()) &&
		Contains(c, r.Right(), r.Right())
}
