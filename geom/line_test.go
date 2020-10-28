package geom

import (
	"fmt"
	"testing"

	"github.com/gravestench/pho/geom/line"
)

func TestBresenhamPoints(t *testing.T) {
	l := line.New(0.33, 0.49, 40, 50)
	points := line.BresenhamPoints(l, 1, nil)

	for _, point := range points {
		fmt.Printf("P(%.2f, %.2f)\n", point.X, point.Y)
	}
}
