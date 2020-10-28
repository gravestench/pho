package triangle

// Offset moves each point (vertex) of a Triangle by a given offset,
// thus moving the entire Triangle by that offset.
func Offset(t *Triangle, ox, oy float64) *Triangle  {
	t.X1 += ox
	t.Y1 += oy
	t.X2 += ox
	t.Y2 += oy
	t.X3 += ox
	t.Y3 += oy

	return t
}
