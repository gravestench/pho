package line

// Extends the start and end points of a Line by the given amounts.
// The amounts can be positive or negative. Positive points will increase the length of the line,
// while negative ones will decrease it.
// Pass a value of zero to leave the start or end point unchanged.
func Extend(l *Line, left, right float64) *Line {
	length := Length(l)
	slopeX, slopeY := l.X2-l.X1, l.Y2-l.Y1

	l.X1 = l.X1 - slopeX / length * left
	l.Y1 = l.Y1 - slopeY / length * left

	l.X2 = l.X2 - slopeX / length * right
	l.Y2 = l.Y2 - slopeY / length * right

	return l
}
