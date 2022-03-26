package complexMatrix

import "testing"

type CopyParams struct {
	height     int
	width      int
	start_real float64
	start_imag float64
	step_real  float64
	step_imag  float64
}

func FuzzCopy(f *testing.F) {
	for _, seed := range []CopyParams{
		{1, 1, 0, 0, 1, 1},
		{2, 2, 1, 0, 3, -1},
	} {
		f.Add(seed.height, seed.width, seed.start_real, seed.start_imag, seed.step_real, seed.step_imag)
	}
	f.Fuzz(func(t *testing.T, height int, width int, start_real float64, start_imag float64, step_real float64, step_imag float64) {
		if height > 0 && width > 0 {
			m := make(immutable, width)
			val := complex(start_real, start_imag)
			step := complex(step_real, step_imag)
			for c := range m {
				m[c] = make([]complex128, height)
				for r := range m[c] {
					m[c][r] = val
					val += step
				}
			}
			o := m.copy()
			if !Equal(o, m) {
				t.Fatalf("Copy failed: %v, %v", m, o)
			}
		}
	})
}
