package line

// PerpendicularSlope calculates the perpendicular slope of the given line.
func PerpendicularSlope(l *Line) float64 {
	return -((l.X2 - l.X1) / (l.Y2 - l.Y1))
}
