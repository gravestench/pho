package point

// CopyFrom copies the values of one Point to a destination Point.
func CopyFrom(source, dest *Point) *Point {
	return dest.SetTo(source.X, source.Y)
}
