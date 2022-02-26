package complexMatrix

import "fmt"

type M interface {
	Dim() (r int, c int)
	Get(r int, c int) complex128
	Set(v complex128, r int, c int) M
	Scale(c complex128) M
	Add(m M) M
	Transpose() M
	Dot(m M) M // A.Dot(B) => AB, B.Dot(A) => BA
	Map(f func(v complex128, r int, c int) complex128) M
	Resize(R int, C int) M
	Immutable() M
	Mutable() M
}

func Real(m M) [][]float64 {
	R, C := m.Dim()
	n := make([][]float64, 0, R)
	for i := 0; i < R; i++ {
		nrow := make([]float64, 0, C)
		for j := 0; j < C; j++ {
			nrow = append(nrow, real(m.Get(i, j)))
		}
		n = append(n, nrow)
	}
	return n
}

func Imag(m M) [][]float64 {
	R, C := m.Dim()
	n := make([][]float64, 0, R)
	for i := 0; i < R; i++ {
		nrow := make([]float64, 0, C)
		for j := 0; j < C; j++ {
			nrow = append(nrow, imag(m.Get(i, j)))
		}
		n = append(n, nrow)
	}
	return n
}

func Equal(a M, b M) bool {
	R, C := a.Dim()
	if bR, bC := b.Dim(); R != bR || C != bC {
		return false
	}
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if a.Get(i, j) != b.Get(i, j) {
				return false
			}
		}
	}
	return true
}

type immutable [][]complex128

func NewImmutable(table [][]complex128) M {
	n := make(immutable, len(table))
	columnChecker := len(table[0])
	for i := range n {
		if columnChecker != len(table[i]) {
			panic("complexMatrix.NewImmutable parameter error: values are not in rectangular shape")
		}
		n[i] = make([]complex128, columnChecker)
		for j := range n[i] {
			n[i][j] = table[i][j]
		}
	}
	return n
}

func combine(r [][]float64, im [][]float64) [][]complex128 {
	R, C := len(r), len(r[0])
	if iR, iC := len(im), len(im[0]); R != iR || C != iC {
		panic("attempting to combine real and imaginary matrixes of different dimensions")
	}
	n := make([][]complex128, 0, R)
	for i := 0; i < R; i++ {
		nrow := make([]complex128, 0, C)
		for j := 0; j < C; j++ {
			nrow = append(nrow, complex(r[i][j], im[i][j]))
		}
		n = append(n, nrow)
	}
	return n
}

func CombineIntoImmutable(real [][]float64, imag [][]float64) M {
	return NewImmutable(combine(real, imag))
}

func (m immutable) copy() immutable {
	if m == nil {
		return nil
	}
	n := make(immutable, len(m))
	copy(n, m)
	return n
}

func (m immutable) Dim() (int, int) {
	if m == nil || len(m) == 0 {
		return 0, 0
	}
	return len(m), len(m[0])
}

func (m immutable) Get(i int, j int) complex128 {
	if m == nil {
		return 0
	}
	return m[i][j]
}

func (m immutable) Set(c complex128, i int, j int) M {
	if R, C := m.Dim(); i < 0 || i >= R || j < 0 || j >= C {
		panic(fmt.Sprintf("complexMatrix.immutable.Set parameter error: i (%d) or j(%d) is out of bounds (%d, %d)", i, j, R, C))
	}
	n := m.copy()
	n[i] = make([]complex128, len(m[0]))
	copy(n[i], m[i])
	n[i][j] = c
	return n
}

func (m immutable) Scale(v complex128) M {
	n := make(immutable, len(m))
	for i := range n {
		n[i] = make([]complex128, len(m[0]))
		for j := range n[i] {
			n[i][j] = m[i][j] * v
		}
	}
	return n
}

func (m immutable) Add(o M) M {
	{
		mR, mC := m.Dim()
		oR, oC := o.Dim()
		if mR != oR || mC != oC {
			panic("dimesion mismatch, can only add matricies of the same dimentions")
		}
	}
	n := make(immutable, len(m))
	for i := range n {
		n[i] = make([]complex128, len(m[0]))
		for j := range n[i] {
			n[i][j] = m.Get(i, j) + o.Get(i, j)
		}
	}
	return n
}

func (m immutable) Transpose() M {
	return transpose{
		wrap: m.copy(),
	}
}

func (m immutable) Dot(B M) M {
	return dot(m, B, m)
}

func (m immutable) Build(d [][]complex128) M {
	return NewImmutable(d)
}

func (m immutable) String() string {
	return SPrintCustom(m, "[", "], ", ", ")
}

func (m immutable) Map(f func(v complex128, r int, c int) complex128) M {
	n := make(immutable, len(m))
	for i := range n {
		n[i] = make([]complex128, len(m[0]))
		for j := range n[i] {
			n[i][j] = f(m.Get(i, j), i, j)
		}
	}
	return n
}

func (m immutable) Resize(R int, C int) M {
	n := make(immutable, R)
	mR, mC := m.Dim()
	for i := range n {
		n[i] = make([]complex128, C)
		for j := range n[i] {
			if i >= mR || j >= mC {
				n[i][j] = 0
			} else {
				n[i][j] = m[i][j]
			}
		}
	}
	return n
}

func (m immutable) Immutable() M {
	return m
}

func (m immutable) Mutable() M {
	return NewMutable(m)
}
