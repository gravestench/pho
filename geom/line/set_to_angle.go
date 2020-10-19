package line

import "math"

// SetToAngle sets a line to a given position, angle and length.
func SetToAngle(l *Line, x, y, angle, length float64) *Line {
	l.X1, l.Y1 = x, y
	l.X2, l.Y2 = x+(math.Cos(angle)*length), y+(math.Sin(angle)*length)

	return l
}
