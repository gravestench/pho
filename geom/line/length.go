package line

import "math"

//  Length calculates the length of the given line.
func Length(l *Line) float64 {
	return math.Sqrt((l.X2 - l.X1) * (l.X2 - l.X1) + (l.Y2 - l.Y1) * (l.Y2 - l.Y1))
}
