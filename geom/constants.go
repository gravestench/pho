package geom

type ShapeType int

const (
	Circle ShapeType = iota
	Ellipse
	Line
	Point
	Polygon
	Rectangle
	Triangle
)