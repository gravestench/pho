package line

import (
	"fmt"
	"testing"
)

func Test_BresenhamPoints(t *testing.T) {
	line := New(0, 0, 4, 5)
	points := BresenhamPoints(line, 1, nil)
	for _, point := range points {
		fmt.Printf("P(%.2f, %.2f)\n", point.X, point.Y)
	}
}
