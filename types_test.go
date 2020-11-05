package complexMatrix

import "testing"

func TestEqual(t *testing.T) {
	vals := make([]*complex128, 0, 10)
	for v := float64(1.0); v <= 10.0; v++ {
		temp := new(complex128)
		*temp = complex(v, v)
		vals = append(vals, temp)
	}
	a := &immutable{
		R: 2,
		C: 2,
		Data: [][]*complex128{
			[]*complex128{vals[0], vals[1]},
			[]*complex128{vals[2], vals[3]},
		},
	}
	if !Equal(a, a) {
		t.Errorf("reflexive equality check failed")
	}

}

func TestImmutable(t *testing.T) {

}
