package tool

import "testing"

func TestInArray(t *testing.T) {
	n := 1
	s := []int{0, 1, 2, 3}
	i, ok := InArray(n, s)
	if !ok {
		t.Errorf("failed to to ok")
	}

	if i != 1 {
		t.Errorf("failed to to get index")
	}
}
