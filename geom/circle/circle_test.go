package circle

import "testing"

func Test_NewCircle(t *testing.T) {
	circle := New(0, 0, 1)

	if circle == nil {
		t.Error("could not create a circle")
	}
}
