package point

import "math"

// GetMagnitude calculates the magnitude of the point,
// which equivalent to the length of the line from the origin to this point.
func GetMagnitude(p *Point) float64 {
	return math.Sqrt(GetMagnitudeSquared(p))
}
