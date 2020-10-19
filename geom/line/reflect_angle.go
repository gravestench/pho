package line

import "github.com/gravestench/pho/phomath"

// ReflectAngle calculates the reflected angle between two lines.
// This is the outgoing angle based on the angle of Line 1 and the normalAngle of Line 2.
func ReflectAngle(a, b *Line) float64 {
	return 2 * NormalAngle(b) - phomath.PI - Angle(a)
}
