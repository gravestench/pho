package line

import "math"

// RotateXY rotates a line around the given coordinates by the given angle in radians.
func RotateAroundXY(l *Line, x, y, angle float64) *Line {
	c, s := math.Cos(angle), math.Sin(angle)

	tx, ty := l.X1-x, l.Y1-y

	l.X1 = tx*c - ty*s + x
	l.Y1 = tx*s + ty*c + y

	tx = l.X2 - x
	ty = l.Y2 - y

	l.X2 = tx*c - ty*s + x
	l.Y2 = tx*s + ty*c + y

	return l
}
