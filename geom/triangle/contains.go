package triangle

func Contains(t *Triangle, x, y float64) bool {
	v0x := t.X3 - t.X1
	v0y := t.Y3 - t.Y1
	v1x := t.X2 - t.X1
	v1y := t.Y2 - t.Y1
	v2x := x - t.X1
	v2y := y - t.Y1

	dot00 := (v0x * v0x) + (v0y * v0y)
	dot01 := (v0x * v1x) + (v0y * v1y)
	dot02 := (v0x * v2x) + (v0y * v2y)
	dot11 := (v1x * v1x) + (v1y * v1y)
	dot12 := (v1x * v2x) + (v1y * v2y)

	// Compute barycentric coordinates
	b := (dot00 * dot11) - (dot01 * dot01)
	inv := 0.

	if b != 0 {
		inv = 1 / b
	}

	u := ((dot11 * dot02) - (dot01 * dot12)) * inv
	v := ((dot00 * dot12) - (dot01 * dot02)) * inv

	return u >= 0 && v >= 0 && (u + v < 1)
}
