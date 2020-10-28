package triangle

import "github.com/gravestench/pho/geom/point"

// GetPoint returns a point on the triangle at a position
func GetPoint(t *Triangle, position float64, out *point.Point) *point.Point {
	if out == nil {
		out = point.New(0, 0)
	}

	l1, l2, l3 := t.GetLineA(), t.GetLineB(), t.GetLineC()

	if position <= 0 || position >= 1 {
		out.X = l1.X1
		out.Y = l1.Y1
		return out
	}

	len1, len2, len3 := l1.Length(), l2.Length(), l3.Length()

	perimeter := len1 + len2 + len3
	p := perimeter * position
	localPosition := 0.

	//  Which line is it on?
	if p < len1 { // Line 1
		localPosition = p / len1
		out.X = l1.X1 + (l1.X2-l1.X1)*(p/len1)
		out.Y = l1.Y1 + (l1.Y2-l1.Y1)*(p/len1)
	} else if p > len1+len2 { // Line 3
		p -= len1 + len2
		localPosition = p / len3
		out.X = l3.X1 + (l3.X2-l3.X1)*localPosition
		out.Y = l3.Y1 + (l3.Y2-l3.Y1)*localPosition
	} else { // Line 2
		p -= len1
		localPosition = p / len2
		out.X = l2.X1 + (l2.X2-l2.X1)*localPosition
		out.Y = l2.Y1 + (l2.Y2-l2.Y1)*localPosition
	}

	return out
}
