package circle

import (
	"github.com/gravestench/pho/geom"
	"github.com/gravestench/pho/geom/point"
	"github.com/gravestench/pho/geom/rectangle"
)

// New creates a new circle
func New(x, y, radius float64) *Circle {
	c := &Circle{
		Type: geom.Circle,
	}

	return c.SetPosition(x, y).SetRadius(radius)
}

type Circle struct {
	Type             geom.ShapeType
	X, Y             float64
	radius, diameter float64
}

// Area returns the circle area
func (c *Circle) Area() float64 {
	return Area(c)
}

// Radius returns the circle radius
func (c *Circle) Radius() float64 {
	return c.radius
}

// SetRadius returns the circle radius, which also updates the diameter
func (c *Circle) SetRadius(r float64) *Circle {
	c.radius = r
	c.diameter = r * 2

	return c
}

// Diameter returns the circle diameter
func (c *Circle) Diameter() float64 {
	return c.diameter
}

// SetDiameter sets the circle diameter, which also updates the radius
func (c *Circle) SetDiameter(d float64) *Circle {
	c.diameter = d
	c.radius = d / 2

	return c
}

// Left returns the left position of the circle
func (c *Circle) Left() float64 {
	return c.X - c.radius
}

// SetLeft sets the left position of the circle, which also sets the x coordinate
func (c *Circle) SetLeft(value float64) *Circle {
	c.X = value + c.radius
	return c
}

func (c *Circle) Right() float64 {
	return c.X + c.radius
}

func (c *Circle) SetRight(value float64) *Circle {
	c.X = value - c.radius
	return c
}

func (c *Circle) Top() float64 {
	return c.Y - c.radius
}

func (c *Circle) SetTop(value float64) *Circle {
	c.Y = value + c.radius
	return c
}

func (c *Circle) Bottom() float64 {
	return c.Y + c.radius
}

func (c *Circle) SetBottom(value float64) *Circle {
	c.Y = value - c.radius
	return c
}

// Contains checks to see if the Circle contains the given x / y coordinates.
func (c *Circle) Contains(x, y float64) bool {
	return Contains(c, x, y)
}

// GetPoint returns a Point object containing the coordinates of a point on the circumference of
// the Circle based on the given angle normalized to the range 0 to 1. I.e. a value of 0.5 will give
// the point at 180 degrees around the circle.
func (c *Circle) GetPoint(position float64, point *point.Point) *point.Point {
	return GetPoint(c, position, point)
}

// GetPoints Returns an array of Point objects containing the coordinates of the points around the
// circumference of the Circle, based on the given quantity or stepRate values.
func (c *Circle) GetPoints(quantity int, stepRate float64, points []*point.Point) []*point.Point {
	return GetPoints(c, quantity, stepRate, points)
}

// Circumference returns the circumference of the given Circle.
func (c *Circle) Circumference() float64 {
	return Circumference(c)
}

// GetRandomPoint returns a uniformly distributed random point from anywhere within the given Circle
func (c *Circle) GetRandomPoint(p *point.Point) *point.Point {
	return GetRandomPoint(c, p)
}

// SetTo sets the x, y and radius of this circle.
func (c *Circle) SetTo(x, y, radius float64) *Circle {
	c.X, c.Y, c.radius = x, y, radius
	c.diameter = 2 * radius

	return c
}

// SetEmpty sets this Circle to be empty with a radius of zero.
func (c *Circle) SetEmpty() *Circle {
	c.radius = 0
	c.diameter = 0

	return c
}

// SetPosition sets the position of this circle
func (c *Circle) SetPosition(x, y float64) *Circle {
	c.X, c.Y = x, y

	return c
}

// IsEmpty checks to see if the Circle is empty: has a radius of zero.
func (c *Circle) IsEmpty() bool {
	return c.radius <= 0
}

// Clone returns a clone of this circle
func (c *Circle) Clone() *Circle {
	return New(c.X, c.Y, c.radius)
}

// Copy the given circle x, y, and radius to this circle.
func (c *Circle) Copy(from *Circle) *Circle {
	return CopyFrom(c, from)
}

// Equals compares the `x`, `y` and `radius` properties of this circle with the other circle.
// Returns `true` if they all match, otherwise returns `false`.
func (c *Circle) Equals(other *Circle) bool {
	return Equals(c, other)
}

// GetBounds returns the bounds (a rectangle) of the Circle object.
func (c *Circle) GetBounds(assignTo *rectangle.Rectangle) *rectangle.Rectangle {
	return GetBounds(c, assignTo)
}

// Offset the circle position by the given x, y
func (c *Circle) Offset(x, y float64) *Circle {
	return Offset(c, x, y)
}

// OffsetPoint offsets the circle with the x,y of the given point
func (c *Circle) OffsetPoint(p *point.Point) *Circle {
	return Offset(c, p.X, p.Y)
}