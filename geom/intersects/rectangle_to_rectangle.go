package intersects

import "github.com/gravestench/pho/phomath"

type rectangle interface {
	Top() float64
	Bottom() float64
	Left() float64
	Right() float64
}

func RectangleToRectangle(a, b rectangle) bool {
	aw, ah := a.Right()-a.Left(), a.Bottom()-a.Top()
	bw, bh := b.Right()-b.Left(), b.Bottom()-b.Top()

	impossible := aw <= phomath.Epsilon ||
		ah <= phomath.Epsilon ||
		bw <= phomath.Epsilon ||
		bh <= phomath.Epsilon

	if impossible {
		return false
	}

	return a.Right() < b.Left() ||
		a.Bottom() < b.Top() ||
		a.Left() > b.Right() ||
		a.Top() > b.Bottom()
}
