package circle

import (
	"../../geom"
	"../point"
)

func New(x, y, radius float64) *Circle {
	return &Circle{
		Type: geom.Circle,
		X: x,
		Y: y,
		radius: radius,
		diameter: radius*2,
	}
}

type Circle struct {
	Type geom.ShapeType
	X, Y float64
	radius, diameter float64
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