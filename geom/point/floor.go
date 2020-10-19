package point

import "math"

// Floor Applies `math.Floor()` to each coordinate of the given Point.
func Floor(p *Point) *Point {
	p.SetTo(math.Floor(p.X), math.Floor(p.Y))
}
