package line

// Offset a line by the given amount.
func Offset(l *Line, x, y float64) *Line {
	l.X1 += x
	l.Y1 += y
	l.X2 += x
	l.Y2 += y

	return l
}
