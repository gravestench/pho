package line

import "math"

// Angle calculates the angle of the line in radians
func Angle(l *Line) float64 {
	return math.Atan2(l.Y2-l.Y1, l.X2-l.X1)
}
