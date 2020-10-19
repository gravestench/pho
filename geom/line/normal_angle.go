package line

import "github.com/gravestench/pho/phomath"

// NormalAngle get the angle of the normal of the given line in radians.
func NormalAngle(l *Line) float64 {
	angle := Angle(l) - phomath.TAU

	return phomath.Wrap(angle, -phomath.PI, phomath.PI)
}
