package circle

import (
	"github.com/gravestench/pho/phomath"
)

func Equals(a, b *Circle) bool {
	dx := (a.X - b.X) * (a.X - b.X)
	dy := (a.Y - b.Y) * (a.Y - b.Y)

	return dx < phomath.Epsilon && dy < phomath.Epsilon
}
