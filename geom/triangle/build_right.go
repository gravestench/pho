package triangle


// BuildRight builds a right triangle, with one 90 degree angle and two acute angles
// The x/y is the coordinate of the 90 degree angle (and will map to x1/y1 in the resulting Triangle)
// w/h can be positive or negative and represent the length of each side
func BuildRight(x, y, width, height float64) *Triangle {
	x1, y1 := x, y
	x2, y2 := x, y-height
	x3, y3 := x+width, y

	return New(x1, y1, x2, y2, x3, y3)
}