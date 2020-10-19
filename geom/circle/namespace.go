package circle

import (
	"github.com/gravestench/pho/geom/point"
	"github.com/gravestench/pho/geom/rectangle"
)

type CircleNamespace interface {
	New(x, y, radius float64) *Circle
	Clone(c *Circle) *Circle
	Contains(c *Circle, x, y float64) bool
	ContainsPoint(c *Circle, p *point.Point) bool
	ContainsRectangle(c *Circle, r *rectangle.Rectangle) bool
	GetPoint(c *Circle, position float64, point *point.Point) *point.Point
	GetPoints(c *Circle, quantity int, stepRate float64, points []*point.Point) []*point.Point
	Circumference(c *Circle) float64
	GetRandomPoint(c *Circle, p *point.Point) *point.Point
	Equals(c *Circle, other *Circle) bool
	GetBounds(c *Circle, assignTo *rectangle.Rectangle) *rectangle.Rectangle
	Offset(c *Circle, x, y float64) *Circle
	OffsetPoint(c *Circle, p *point.Point) *Circle
}

type Namespace struct{}

func (n *Namespace) New(x, y, radius float64) *Circle {
	return New(x, y, radius)
}

// Clone returns a clone of this circle
func (n *Namespace) Clone(c *Circle) *Circle {
	return Clone(c)
}

// Contains checks to see if the Circle contains the given x / y coordinates.
func (n *Namespace) Contains(c *Circle, x, y float64) bool {
	return Contains(c, x, y)
}

// ContainsPoint checks to see if the Circle contains the given point.
func (n *Namespace) ContainsPoint(c *Circle, p *point.Point) bool {
	return ContainsPoint(c, p)
}

// ContainsRectangle checks to see if the Circle contains the given rectangle.
func (n *Namespace) ContainsRectangle(c *Circle, r *rectangle.Rectangle) bool {
	return ContainsRectangle(c, r)
}

// GetPoint returns a Point object containing the coordinates of a point on the circumference of
// the Circle based on the given angle normalized to the range 0 to 1. I.e. a value of 0.5 will give
// the point at 180 degrees around the circle.
func (n *Namespace) GetPoint(c *Circle, position float64, point *point.Point) *point.Point {
	return GetPoint(c, position, point)
}

// GetPoints Returns an array of Point objects containing the coordinates of the points around the
// circumference of the Circle, based on the given quantity or stepRate values.
func (n *Namespace) GetPoints(c *Circle, quantity int, stepRate float64,
	points []*point.Point) []*point.Point {
	return GetPoints(c, quantity, stepRate, points)
}

// Circumference returns the circumference of the given Circle.
func (n *Namespace) Circumference(c *Circle) float64 {
	return Circumference(c)
}

// GetRandomPoint returns a uniformly distributed random point from anywhere within the given Circle
func (n *Namespace) GetRandomPoint(c *Circle, p *point.Point) *point.Point {
	return GetRandomPoint(c, p)
}

// Equals compares the `x`, `y` and `radius` properties of this circle with the other circle.
// Returns `true` if they all match, otherwise returns `false`.
func (n *Namespace) Equals(c *Circle, other *Circle) bool {
	return Equals(c, other)
}

// GetBounds returns the bounds (a rectangle) of the Circle object.
func (n *Namespace) GetBounds(c *Circle, assignTo *rectangle.Rectangle) *rectangle.Rectangle {
	return GetBounds(c, assignTo)
}

// Offset the circle position by the given x, y
func (n *Namespace) Offset(c *Circle, x, y float64) *Circle {
	return Offset(c, x, y)
}

// OffsetPoint offsets the circle with the x,y of the given point
func (n *Namespace) OffsetPoint(c *Circle, p *point.Point) *Circle {
	return Offset(c, p.X, p.Y)
}
