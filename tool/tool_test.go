package tool

import "testing"

func TestInArray(t *testing.T) {
	assertEqual := func(exp interface{}, val interface{}) {
		if val != exp {
			t.Errorf("Expected %v, got %v.", exp, val)
		}
	}

	i, ok := InArray(1, []int{0, 1, 2, 3})
	assertEqual(true, ok)
	assertEqual(1, i)

	i, ok = InArray(4, []int{0, 1, 2, 3})
	assertEqual(false, ok)
	assertEqual(-1, i)
}
