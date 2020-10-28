package triangle

import "github.com/gravestench/pho/geom/point"

// BuildFromPolyon takes an array of vertex coordinates, and optionally an array of hole indices,
// then returns an array of Triangle instances, where the given vertices have been decomposed
// into a series of triangles.
func BuildFromPolygon(coords, holes []float64, sx, sy float64, out []*point.Point) []*point.Point {
	panic("BuildFromPolyon not implemented yet!")
}
