package line

import "math"

// Height calculates the height of the given line.
func Height(l *Line) float64 {
	return math.Abs(l.Y1 - l.Y2)
}
