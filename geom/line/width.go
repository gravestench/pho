package line

import "math"

// Width calculates the width of the given line.
func Width(l *Line) float64 {
	return math.Abs(l.X1 - l.X2)
}
