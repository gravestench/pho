package line

import (
	"fmt"
	"testing"
)

func TestBresenhamPoints(t *testing.T) {
	line := New(0.33, 0.49, 40, 50)
	points := BresenhamPoints(line, 1, nil)

	for _, point := range points {
		fmt.Printf("P(%.2f, %.2f)\n", point.X, point.Y)
	}
}
