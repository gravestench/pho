package point

import "math"

// Ceil Applies `math.Ceil()` to each coordinate of the given Point.
func Ceil(p *Point) *Point {
	p.SetTo(math.Ceil(p.X), math.Ceil(p.Y))
}
