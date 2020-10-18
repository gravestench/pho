package circle

import (
	"github.com/gravestench/pho/geom/point"
	"github.com/gravestench/pho/phomath"
	"math"
	"math/rand"
)

// GetRandomPoint returns a uniformly distributed random point from anywhere within the given Circle
func GetRandomPoint(c *Circle, p *point.Point) *point.Point {
	if p == nil {
		p = point.New(0, 0)
	}

	t := phomath.PI2 * rand.Float64()
	u := rand.Float64() * 2 - 1 // -1..1
	x := u * math.Cos(t)
	y := u * math.Sin(t)

	p.X = c.X + (x * c.radius)
	p.Y = c.Y + (y * c.radius)

	return p
}