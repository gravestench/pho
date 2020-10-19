package line

import (
	"github.com/gravestench/pho/geom/point"
)

type LineNamespace interface {
	New(x1, y1, x2, y2 float64) *Line
	BresenhamPoints(l *Line, stepRate int, out []*point.Point) []*point.Point
	Length(l *Line) float64
	GetPoint(l *Line, position float64, out *point.Point) *point.Point
	GetPoints(l *Line, quantity int, stepRate float64, out []*point.Point) []*point.Point
	GetRandomPoint(l *Line, assignTo *point.Point) *point.Point
	Angle(l *Line) float64
	CenterOn(l *Line, x, y float64) *Line
	Clone(l *Line) *Line
	CopyFrom(l *Line, from *Line) *Line
	Equals(l *Line, other *Line) bool
	GetMidPoint(l *Line, out *point.Point) *point.Point
	GetNearestPoint(l *Line, p, out *point.Point) *point.Point
	GetNormal(l *Line, p, out *point.Point) *point.Point
	GetShortestDistance(l *Line, p *point.Point) float64
	Height(l *Line) float64
	Width(l *Line) float64
	NormalAngle(l *Line) float64
	NormalX(l *Line) float64
	NormalY(l *Line) float64
	Offset(l *Line, x, y float64) *Line
	PerpendicularSlope(l *Line) float64
	ReflectAngle(l *Line, other *Line) float64
	Rotate(l *Line, angle float64) *Line
	RotateAroundPoint(l *Line, p *point.Point, angle float64) *Line
	RotateAroundXY(l *Line, x, y, angle float64) *Line
	SetToAngle(l *Line, x, y, angle, length float64) *Line
	Slope(l *Line) float64
}

type Namespace struct{}

// New creates a new line
func (*Namespace) New(x1, y1, x2, y2 float64) *Line {
	return New(x1, y1, x2, y2)
}

// BresenhamPoints uses Bresenham's line algorithm to return an array of all coordinates
// on this line.
func (*Namespace) BresenhamPoints(l *Line, stepRate int, out []*point.Point) []*point.Point {
	return BresenhamPoints(l, stepRate, out)
}

// Length calculates the length of the line.
func (*Namespace) Length(l *Line) float64 {
	return Length(l)
}

// GetPoint on a line that's a given percentage along its length.
func (*Namespace) GetPoint(l *Line, position float64, out *point.Point) *point.Point {
	return GetPoint(l, position, out)
}

// Get a number of points along a line's length.
// Provide a `quantity` to get an exact number of points along the line.
func (*Namespace) GetPoints(l *Line, quantity int, stepRate float64, out []*point.Point) []*point.Point {
	return GetPoints(l, quantity, stepRate, out)
}

// GetRandomPoint picks a random point along the line and assigns it to the argument point.
// If argument is nil, a new point is created and returned.
func (*Namespace) GetRandomPoint(l *Line, assignTo *point.Point) *point.Point {
	return GetRandomPoint(l, assignTo)
}


// Angle calculates the angle of the line in radians
func (*Namespace) Angle(l *Line, ) float64 {
	return Angle(l)
}

// CenterOn centers a line on the given coordinates.
func (*Namespace) CenterOn(l *Line, x, y float64) *Line {
	return CenterOn(l, x, y)
}

// Clone creates a clone of the line
func (*Namespace) Clone(l *Line, ) *Line {
	return Clone(l)
}

// CopyFrom copies the values of one line to the line.
func (*Namespace) CopyFrom(l *Line, from *Line) *Line {
	return CopyFrom(from, l)
}

// Equals checks if the line is approximately equal to another
func (*Namespace) Equals(l *Line, other *Line) bool {
	return Equals(l, other)
}

// GetMidPoint get the midpoint of the line.
// Assigns to the `out` point, or creates one if nil.
func (*Namespace) GetMidPoint(l *Line, out *point.Point) *point.Point {
	return GetMidPoint(l, out)
}

// GetNearestPoint get the nearest point on a line perpendicular to the given point.
func (*Namespace) GetNearestPoint(l *Line, p, out *point.Point) *point.Point {
	return GetNearestPoint(l, p, out)
}

// GetNormal calculates the normal of the line.
// The normal of a line is a vector that points perpendicular from it.
func (*Namespace) GetNormal(l *Line, p, out *point.Point) *point.Point {
	return GetNormal(l, p, out)
}

// GetShortestDistance get the shortest distance from a Line to the given Point.
func (*Namespace) GetShortestDistance(l *Line, p *point.Point) float64 {
	return GetShortestDistance(l, p)
}

// Height calculates the height of the line.
func (*Namespace) Height(l *Line, ) float64 {
	return Height(l)
}

// Width calculates the width of the given line.
func (*Namespace) Width(l *Line, ) float64 {
	return Width(l)
}

// NormalAngle get the angle of the normal of the line in radians.
func (*Namespace) NormalAngle(l *Line, ) float64 {
	return NormalAngle(l)
}

// NormalX returns the x component of the normal vector of the line.
// The normal of a line is a vector that points perpendicular from it.
func (*Namespace) NormalX(l *Line, ) float64 {
	return NormalX(l)
}

// NormalX returns the x component of the normal vector of the line.
// The normal of a line is a vector that points perpendicular from it.
func (*Namespace) NormalY(l *Line, ) float64 {
	return NormalY(l)
}

// Offset the line by the given amount.
func (*Namespace) Offset(l *Line, x, y float64) *Line {
	return Offset(l, x, y)
}

// PerpendicularSlope calculates the perpendicular slope of the line.
func (*Namespace) PerpendicularSlope(l *Line, ) float64 {
	return PerpendicularSlope(l)
}

// ReflectAngle calculates the reflected angle between the line and another line.
// This is the outgoing angle based on the angle of Line 1 and the normalAngle of Line 2.
func (*Namespace) ReflectAngle(l *Line, other *Line) float64 {
	return ReflectAngle(l, other)
}

// Rotate a line around its midpoint by the given angle in radians.
func (*Namespace) Rotate(l *Line, angle float64) *Line {
	return Rotate(l, angle)
}

// RotateAroundPoint rotates a line around a point by the given angle in radians.
func (*Namespace) RotateAroundPoint(l *Line, p *point.Point, angle float64) *Line {
	return RotateAroundPoint(l, p, angle)
}

// RotateXY rotates a line around the given coordinates by the given angle in radians.
func (*Namespace) RotateAroundXY(l *Line, x, y, angle float64) *Line {
	return RotateAroundXY(l, x, y, angle)
}

// SetToAngle sets a line to a given position, angle and length.
func (*Namespace) SetToAngle(l *Line, x, y, angle, length float64) *Line {
	return SetToAngle(l, x, y, angle, length)
}

// Slope calculates the slope of the line.
func (*Namespace) Slope(l *Line, ) float64 {
	return Slope(l)
}
