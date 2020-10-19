package point

import "github.com/gravestench/pho/geom/rectangle"

// GetRectangleFromPoints calculates the Axis Aligned Bounding Box ( or aabb) from an array of
// points.
func GetRectangleFromPoints(points []*Point) *rectangle.Rectangle {
	return rectangle.New(0, 0, 0, 0).MergePoints(points)
}
