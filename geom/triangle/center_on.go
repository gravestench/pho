package triangle

import "github.com/gravestench/pho/geom/point"

// CenterOn positions the Triangle so that it is centered on the given coordinates.
// Argument `fn` is the centering function.
func CenterOn(t *Triangle, x, y float64, fn func(*Triangle, *point.Point) *point.Point) *Triangle {
	if fn == nil {
		fn = Centroid
	}

	center := fn(t, nil)
	dx, dy := x-center.X, y-center.Y

	return Offset(t, dx, dy)
}
