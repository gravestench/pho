package phomath

// FromPercent returns a value based on the range between `min` and `max` and the percentage given.
func FromPercent(percent, min, max float64) float64 {
	percent = Clamp(percent, 0, 1)

	return (max - min) * percent
}
