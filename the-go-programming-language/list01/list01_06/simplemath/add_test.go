package simplemath

import "testing"
import "../../list01_06/simplemath"

func TestAdd1(t *testing.T) {
	r := simplemath.Add(1, 2)
	if r != 3 {
		t.Errorf("Add(1,2) failed,Got %d, expected 3.", r)
	}
}
