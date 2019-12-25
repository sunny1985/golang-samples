package simplemath

import "testing"
import "../../list01_08/simplemath"

func TestSqrt1(t *testing.T)  {
	v := simplemath.Sqrt(16)
	if v != 4 {
		t.Errorf("Sqrt(16) failed. Got %v,expected 4.", v)
	}
}
