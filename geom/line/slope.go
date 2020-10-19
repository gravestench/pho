package line

// Slope calculates the slope of the given line.
func Slope(l *Line) float64 {
	return (l.Y2 - l.Y1) / (l.X2 - l.X1)
}
