package triangle

import (
	"github.com/gravestench/pho/geom"
	"github.com/gravestench/pho/geom/line"
	"github.com/gravestench/pho/geom/point"
)

// New creates a new triangle
func New(x1, y1, x2, y2, x3, y3 float64) *Triangle {
	return &Triangle{
		Type: geom.Triangle,
		X1:   x1,
		Y1:   y1,
		X2:   x2,
		Y2:   y2,
		X3:   x3,
		Y3:   y3,
	}
}

// Triangle defines a Triangle segment, a part of a triangle between two endpoints.
type Triangle struct {
	Type   geom.ShapeType
	X1, Y1 float64
	X2, Y2 float64
	X3, Y3 float64
}

// Contains
func (t *Triangle) Contains(x, y float64) bool {
	return Contains(t, x, y)
}

// GetPoint
func (t *Triangle) GetPoint(position float64, out *point.Point) *point.Point {
	return GetPoint(t, position, out)
}

// GetPoints
func (t *Triangle) GetPoints() []*point.Point {

}

// GetRandomPoint
func (t *Triangle) GetRandomPoint() *point.Point {

}

// SetTo
func (t *Triangle) SetTo(x1, y1, x2, y2, x3, y3 float64) *Triangle {

}

// GetLineA
func (t *Triangle) GetLineA() *line.Line {

}

// GetLineB
func (t *Triangle) GetLineB() *line.Line {

}

// GetLineC
func (t *Triangle) GetLineC() *line.Line {

}

// Area calculates the area of the given rectangle.
func (t *Triangle) Area() float64 {
	return t.Width * t.Height
}

// Left returns the left position of the rectangle
func (t *Triangle) Left() float64 {
	return t.X
}

// SetLeft sets the left position of the rectangle, which also sets the x coordinate
func (t *Triangle) SetLeft(value float64) *Triangle {
	if value >= t.Right() {
		t.Width = 0
	} else {
		t.Width = t.Right() - value
	}

	t.X = value
	return r
}

// Right returns the right position of the rectangle
func (t *Triangle) Right() float64 {
	return t.X + t.Width
}

// SetRight sets the right position of the rectangle, which adjusts the width
func (t *Triangle) SetRight(value float64) *Triangle {
	if value <= t.X {
		t.Width = 0
	} else {
		t.Width = value - t.X
	}

	t.X = value

	return r
}

// Top returns the top position of the rectangle
func (t *Triangle) Top() float64 {
	return t.Y
}

// SetTop sets the top position of the rectangle, which also adjusts the Y coordinate
func (t *Triangle) SetTop(value float64) *Triangle {
	if value >= t.Bottom() {
		t.Height = 0
	} else {
		t.Height = t.Bottom() - value
	}

	return r
}

// Bottom returns the bottom position of the rectangle
func (t *Triangle) Bottom() float64 {
	return t.Y + t.Height
}

// SetBottom sets the bottom position of the rectangle, which also adjusts the height
func (t *Triangle) SetBottom(value float64) *Triangle {
	if value <= t.Y {
		t.Height = 0
	} else {
		t.Height = value - t.Y
	}

	return r
}

func (t *Triangle) CenterX() float64 {
	return t.X + t.Width/2
}

func (t *Triangle) SetCenterX(value float64) *Triangle {
	t.X = value - t.Width/2
	return r
}

func (t *Triangle) CenterY() float64 {
	return t.Y + t.Height/2
}

func (t *Triangle) SetCenterY(value float64) *Triangle {
	t.Y = value - t.Height/2

	return r
}

// CenterOn moves the top-left corner of a Rectangle so that its center is at the given coordinates.
func (t *Triangle) CenterOn(x, y float64) *Triangle {
	return CenterOn(r, x, y)
}







