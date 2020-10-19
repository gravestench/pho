package point

// GetMagnitudeSquared calculates the magnitude squared of the point.
func GetMagnitudeSquared(p *Point) float64 {
	return (p.X * p.X) + (p.Y * p.Y)
}
