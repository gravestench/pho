package circle

// Copies the `x`, `y` and `radius` properties from the `source` Circle
// into the given `dest` Circle, then returns the `dest` Circle.
func CopyFrom(source, dest *Circle) *Circle {
	return dest.SetTo(source.X, source.Y, source.radius)
}
