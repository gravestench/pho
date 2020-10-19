package line

import "github.com/gravestench/pho/geom/point"

// RotateAroundPoint rotates a line around a point by the given angle in radians.
func RotateAroundPoint(l *Line, p *point.Point, angle float64) *Line {
	return RotateAroundXY(l, p.X, p.Y, angle)
}
