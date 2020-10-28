package geom

import (
	"github.com/gravestench/pho/geom/circle"
	"testing"
)

func Test_NewCircle(t *testing.T) {
	c := circle.New(0, 0, 1)

	if c == nil {
		t.Error("could not create a circle")
	}
}
