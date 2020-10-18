package rectangle

import (
	"github.com/gravestench/pho/phomath"
	"math"
)

// GetAspectRatio returns the aspect ratio (width/height) of the given rectangle
func GetAspectRatio(r *Rectangle) float64 {
	if r.Height < phomath.Epsilon {
		return math.NaN()
	}

	return  r.Width / r.Height
}
