package line

// Clone creates a clone of a line
func Clone(l *Line) *Line {
	return New(l.X1, l.Y1, l.X2, l.Y2)
}
