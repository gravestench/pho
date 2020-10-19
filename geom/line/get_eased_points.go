package line

import (
	"github.com/gravestench/pho/geom/point"
	"github.com/gravestench/pho/phomath/distance"
	"github.com/gravestench/pho/tweens/builders"
)

// GetEasedPoints returns an array of `quantity` Points where each point is taken from the given
// Line, spaced out according to the ease function specified.
//
// ```golang
// 	line := new Pho.Geom.Line(100, 300, 700, 300)
// 	points := Pho.Geom.Line.GetEasedPoints(line, 'Sine.easeOut', 32, 0, nil)
// ```
//
// In the above example, the `points` array will contain 32 points spread-out across
// the length of `line`, where the position of each point is determined by the `Sine.out`
// ease function.
//
// You can optionally provide a min distance. In this case, the resulting points
// are checked against each other, and if they are `< minDistance` distance apart,
// they are dropped from the result. This can help avoid lots of clustered points at
// far ends of the line with tightly-packed eases such as Quartic. Leave the value set
// to zero to skip this check.
//
// Note that if you provide a collinear threshold, the resulting array may not always
// contain `quantity` points.
func GetEasedPoints(
	l *Line,
	ease interface{},
	quantity int,
	minDistance float64,
	easeParams []float64,
) []*point.Point {
	result := make([]*point.Point, 0)
	x1, y1 := l.X1, l.Y1
	spaceX, spaceY := l.X2-l.X1, l.Y2-l.Y1

	easeFn := builders.GetEaseFunction(ease, easeParams)

	v := float64(0)
	q := quantity - 1

	for idx := 0; idx < q; idx++ {
		v = easeFn(float64(idx) / float64(q))
		result = append(result, point.New(x1+(spaceX*v), y1+(spaceY*v)))
	}

	//  Always include the end of the line
	v = easeFn(1)
	result = append(result, point.New(x1+(spaceX*v), y1+(spaceY*v)))

	//  Remove collinear parts
	if minDistance > 0 {
		prevPoint := result[0]

		//  Store the new result here
		sortedResults := append(make([]*point.Point, 0), prevPoint)

		for idx := 1; idx < len(result) - 1; idx++ {
			var point = result[idx]

			if distance.BetweenPoints(prevPoint, point) >= minDistance {
				sortedResults = append(sortedResults, point)
				prevPoint = point
			}
		}

		//  Top and tail
		var endPoint = result[len(result) - 1]

		if distance.BetweenPoints(prevPoint, endPoint) < minDistance {
			sortedResults = sortedResults[:len(sortedResults)-1]
		}

		sortedResults = append(sortedResults, endPoint)

		return sortedResults
	} else {
		return result
	}
}
