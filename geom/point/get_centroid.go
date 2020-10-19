package point

// GetCentroid gets the centroid or geometric center of a plane figure ( the arithmetic mean
// position of all the points in the figure). Informally, it is the point at which a cutout of the
// shape could be perfectly balanced on the tip of a pin.
func GetCentroid(points []*Point, out *Point) *Point {
	if out == nil {
		out = New(0, 0)
	}

	if points == nil {
		points = make([]*Point, 0)
	}

	numPoints := len(points)

	if numPoints < 1 {
		return out
	} else if numPoints == 1 {
		out.X, out.Y = points[0].X, points[1].X
	} else {
		for idx := 0; idx < numPoints; idx++ {
			out.X += points[idx].X
			out.Y += points[idx].Y
		}

		out.X /= float64(numPoints);
		out.Y /= float64(numPoints);
	}

	return out
}
