package line

import (
	"math"

	"github.com/gravestench/pho/geom"
	"github.com/gravestench/pho/geom/point"
)

// New creates a new line
func New(x1, y1, x2, y2 float64) *Line {
	return &Line{
		Type: geom.Line,
		X1:   x1,
		Y1:   y1,
		X2:   x2,
		Y2:   y2,
	}
}

// Line defines a Line segment, a part of a line between two endpoints.
type Line struct {
	Type   geom.ShapeType
	X1, Y1 float64
	X2, Y2 float64
}

// Length calculates the length of the line.
func (l *Line) Length() float64 {
	return Length(l)
}

// Left returns the left position of the line
func (l *Line) Left() float64 {
	return math.Min(l.X1, l.X2)
}

// SetLeft sets the left position of the line
func (l *Line) SetLeft(value float64) *Line {
	if l.X1 < l.X2 {
		l.X1 = value
	} else {
		l.X2 = value
	}

	return l
}

// Right returns the right position of the line
func (l *Line) Right() float64 {
	return math.Max(l.X1, l.X2)
}

// SetRight sets the right position of the line
func (l *Line) SetRight(value float64) *Line {
	if l.X1 > l.X2 {
		l.X1 = value
	} else {
		l.X2 = value
	}

	return l
}

// Top returns the top position of the line
func (l *Line) Top() float64 {
	return math.Min(l.Y1, l.Y2)
}

// SetTop sets the top position of the line
func (l *Line) SetTop(value float64) *Line {
	if l.Y1 < l.Y2 {
		l.Y1 = value
	} else {
		l.Y2 = value
	}

	return l
}

// Bottom returns the bottom position of the line
func (l *Line) Bottom() float64 {
	return math.Max(l.Y1, l.Y2)
}

// SetBottom sets the bottom position of the line
func (l *Line) SetBottom(value float64) *Line {
	if l.Y1 > l.Y2 {
		l.Y1 = value
	} else {
		l.Y2 = value
	}

	return l
}

// GetPoint on a line that's a given percentage along its length.
func (l *Line) GetPoint(position float64, out *point.Point) *point.Point {
	return GetPoint(l, position, out)
}

// Get a number of points along a line's length.
// Provide a `quantity` to get an exact number of points along the line.
func (l *Line) GetPoints(quantity int, stepRate float64, out []*point.Point) []*point.Point {
	return GetPoints(l, quantity, stepRate, out)
}

// GetRandomPoint picks a random point along the line and assigns it to the argument point.
// If argument is nil, a new point is created and returned.
func (l *Line) GetRandomPoint(assignTo *point.Point) *point.Point {
	return GetRandomPoint(l, assignTo)
}

// SetTo sets new coordinates for the line endpoints.
func (l *Line) SetTo(x1, y1, x2, y2 float64) *Line {
	l.X1, l.Y1, l.X2, l.Y2 = x1, y1, x2, y2
	return l
}

// Returns a point that corresponds to the start of the line segment.
func (l *Line) GetPointA(assignTo *point.Point) *point.Point {
	if assignTo == nil {
		assignTo = point.New(0, 0)
	}

	return assignTo.SetTo(l.X1, l.Y1)
}

// Returns a point that corresponds to the end of the line segment.
func (l *Line) GetPointB(assignTo *point.Point) *point.Point {
	if assignTo == nil {
		assignTo = point.New(0, 0)
	}

	return assignTo.SetTo(l.X2, l.Y2)
}

// Angle calculates the angle of the line in radians
func (l *Line) Angle() float64 {
	return Angle(l)
}

// CenterOn centers a line on the given coordinates.
func (l *Line) CenterOn(x, y float64) *Line {
	return CenterOn(l, x, y)
}

// Clone creates a clone of the line
func (l *Line) Clone() *Line {
	return Clone(l)
}

// CopyFrom copies the values of one line to the line.
func (l *Line) CopyFrom(from *Line) *Line {
	return CopyFrom(from, l)
}

// Equals checks if the line is approximately equal to another
func (l *Line) Equals(other *Line) bool {
	return Equals(l, other)
}

// GetMidPoint get the midpoint of the line.
// Assigns to the `out` point, or creates one if nil.
func (l *Line) GetMidPoint(out *point.Point) *point.Point {
	return GetMidPoint(l, out)
}

// GetNearestPoint get the nearest point on a line perpendicular to the given point.
func (l *Line) GetNearestPoint(p, out *point.Point) *point.Point {
	return GetNearestPoint(l, p, out)
}

// GetNormal calculates the normal of the line.
// The normal of a line is a vector that points perpendicular from it.
func (l *Line) GetNormal(p, out *point.Point) *point.Point {
	return GetNormal(l, p, out)
}

// GetShortestDistance get the shortest distance from a Line to the given Point.
func (l *Line) GetShortestDistance(p *point.Point) float64 {
	return GetShortestDistance(l, p)
}

// Height calculates the height of the line.
func (l *Line) Height() float64 {
	return Height(l)
}

// Width calculates the width of the given line.
func (l *Line) Width() float64 {
	return Width(l)
}

// NormalAngle get the angle of the normal of the line in radians.
func (l *Line) NormalAngle() float64 {
	return NormalAngle(l)
}

// NormalX returns the x component of the normal vector of the line.
// The normal of a line is a vector that points perpendicular from it.
func (l *Line) NormalX() float64 {
	return NormalX(l)
}

// NormalX returns the x component of the normal vector of the line.
// The normal of a line is a vector that points perpendicular from it.
func (l *Line) NormalY() float64 {
	return NormalY(l)
}

// Offset the line by the given amount.
func (l *Line) Offset(x, y float64) *Line {
	return Offset(l, x, y)
}

// PerpendicularSlope calculates the perpendicular slope of the line.
func (l *Line) PerpendicularSlope() float64 {
	return PerpendicularSlope(l)
}

// ReflectAngle calculates the reflected angle between the line and another line.
// This is the outgoing angle based on the angle of Line 1 and the normalAngle of Line 2.
func (l *Line) ReflectAngle(other *Line) float64 {
	return ReflectAngle(l, other)
}

// Rotate a line around its midpoint by the given angle in radians.
func (l *Line) Rotate(angle float64) *Line {
	return Rotate(l, angle)
}

// RotateAroundPoint rotates a line around a point by the given angle in radians.
func (l *Line) RotateAroundPoint(p *point.Point, angle float64) *Line {
	return RotateAroundPoint(l, p, angle)
}

// RotateXY rotates a line around the given coordinates by the given angle in radians.
func (l *Line) RotateAroundXY(x, y, angle float64) *Line {
	return RotateAroundXY(l, x, y, angle)
}

// SetToAngle sets a line to a given position, angle and length.
func (l *Line) SetToAngle(x, y, angle, length float64) *Line {
	return SetToAngle(l, x, y, angle, length)
}

// Slope calculates the slope of the line.
func (l *Line) Slope() float64 {
	return Slope(l)
}