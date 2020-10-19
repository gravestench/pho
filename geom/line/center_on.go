package line

// CenterOn centers a line on the given coordinates.
func CenterOn(l *Line, x, y float64) *Line {
	tx, ty := x-((l.X1+l.X2)/2), y-((l.Y1+l.Y2)/2)

	l.X1 += tx
	l.Y1 += ty

	l.X2 += tx
	l.Y2 += ty

	return l
}
