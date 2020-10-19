package line

// Rotate a line around its midpoint by the given angle in radians.
func Rotate(l *Line, angle float64) *Line {
	var x = (l.X1 + l.X2) / 2;
	var y = (l.Y1 + l.Y2) / 2;

	return RotateAroundXY(l, x, y, angle);
}