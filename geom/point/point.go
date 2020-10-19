package point

import (
	"github.com/gravestench/pho/geom"
)

// New creates a new point
func New(x, y float64) *Point {
	return &Point{
		Type: geom.Point,
		X: x,
		Y: y,
	}
}

// Point defines a Point in 2D space, with an x and y component.
type Point struct {
	Type geom.ShapeType
	X, Y float64
}

// SetTo sets the x and y coordinates of the point to the given values.
func (p *Point) SetTo(x, y float64) *Point {
	p.X, p.Y = x, y
	return p
}

// Ceil Applies `math.Ceil()` to each coordinate of the given Point.
func (p *Point) Ceil() *Point {
	return Ceil(p)
}

// Floor Applies `math.Floor()` to each coordinate of the given Point.
func (p *Point) Floor() *Point {
	return Floor(p)
}

// Clone the given point.
func (p *Point) Clone() *Point {
	return Clone(p)
}

// CopyFrom copies the values of one Point to a destination Point.
func (p *Point) CopyFrom(source *Point) *Point {
	return CopyFrom(source, p)
}

// Equals compares two `Point` objects to see if they are equal.
func (p *Point) Equals(other *Point) bool {
	return Equals(p, other)
}

// GetMagnitude calculates the magnitude of the point,
// which equivalent to the length of the line from the origin to this point.
func (p *Point) GetMagnitude() float64 {
	return GetMagnitude(p)
}

// GetMagnitudeSquared calculates the magnitude squared of the point.
func (p *Point) GetMagnitudeSquared() float64 {
	return GetMagnitudeSquared(p)
}

// Interpolate returns the linear interpolation point between this and another point, based on `t`.
func (p *Point) Interpolate(other *Point, t float64, out *Point) *Point {
	return Interpolate(p, other, t, out)
}

// Invert swaps the X and the Y coordinate of a point.
func (p *Point) Invert() *Point {
	return Invert(p)
}

// Negative flips the sign of the X and the Y coordinate of a point.
func (p *Point) Negative() *Point {
	return Negative(p)
}

// Project calculates the vector projection of `pointA` onto the nonzero `pointB`. This is the
// orthogonal projection of `pointA` onto a straight line parallel to `pointB`.
func (p *Point) Project(other, out *Point) *Point {
	return Project(p, other, out)
}

// SetMagnitude calculates the magnitude of the point,
// which equivalent to the length of the line from the origin to this point.
func (p *Point) SetMagnitude(magnitude float64) *Point {
	return SetMagnitude(p, magnitude)
}