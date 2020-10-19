package line

// CopyFrom copies the values of one line to another line.
func CopyFrom(from, to *Line) *Line {
	return  to.SetTo(from.X1, from.Y1, from.X1, from.Y2)
}
