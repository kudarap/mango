package tool

import "testing"

func TestInArray(t *testing.T) {
	i, ok := InArray(1, []int{0, 1, 2, 3})
	if !ok {
		t.Errorf("failed to find needle")
	} else if i != 1 {
		t.Errorf("failed to to get needle index")
	}
}

func TestNotInArray(t *testing.T) {
	i, ok := InArray(4, []int{0, 1, 2, 3})
	if ok {
		t.Errorf("it should not be found")
	} else if i != -1 {
		t.Errorf("it should not have index")
	}
}
