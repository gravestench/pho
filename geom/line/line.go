package line

import (
	"github.com/gravestench/pho/geom"
	"github.com/gravestench/pho/geom/point"
	"math"
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