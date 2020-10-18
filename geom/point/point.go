package point

import (
	"../../geom"
)

func New(x, y float64) *Point {
	return &Point{
		Type: geom.Point,
		X: x,
		Y: y,
	}
}

type Point struct {
	Type geom.ShapeType
	X, Y float64
}

func (p *Point) SetTo(x, y float64) *Point {
	p.X, p.Y = x, y
	return p
}
