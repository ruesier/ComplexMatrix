package complexMatrix

type M interface {
	Dim() (r int, c int)
	Get(r int, c int) complex128
	Set(v complex128, r int, c int) M
	Scale(c complex128) M
	Add(m M) M
	Transpose() M
	Dot(m M) M // A.Dot(B) => AB, B.Dot(A) => BA
	// Map(f func(v complex128, r int, c int) complex128) M
	// Resize(R int, C int) M
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

// TODO: loose and strict implementations of M

type immutable struct {
	R    int
	C    int
	Data [][]*complex128
}

func NewImmutable(table [][]complex128) M {
	n := new(immutable)
	n.R = len(table)
	if n.R < 1 {
		return &immutable{
			R: 0,
			C: 0,
		}
	}
	n.Data = make([][]*complex128, n.R)
	n.C = len(table[0])
	if n.C < 1 {
		return n
	}
	for i := range n.Data {
		n.Data[i] = make([]*complex128, n.C)
		if len(table[i]) != n.C {
			panic("Matrix should be rectangular")
		}
		for j := range n.Data[i] {
			n.Data[i][j] = &table[i][j]
		}
	}
	return n
}

func (m *immutable) copy() *immutable {
	if m == nil {
		return nil
	}
	n := &immutable{
		R:    m.R,
		C:    m.C,
		Data: make([][]*complex128, len(m.Data)),
	}
	copy(n.Data, m.Data)
	return n
}

func (m *immutable) Dim() (int, int) {
	if m == nil {
		return 0, 0
	}
	return m.R, m.C
}

func (m *immutable) Get(i int, j int) complex128 {
	return *(m.Data[i][j])
}

func (m *immutable) Set(c complex128, i int, j int) M {
	n := m.copy()
	n.Data[i] = make([]*complex128, m.C)
	copy(n.Data[i], m.Data[i])
	n.Data[i][j] = &c
	return n
}

func (m *immutable) Scale(v complex128) M {
	n := new(immutable)
	n.R = m.R
	n.C = m.C
	n.Data = make([][]*complex128, n.R)
	for i := range n.Data {
		n.Data[i] = make([]*complex128, n.C)
		for j := range n.Data[i] {
			c := new(complex128)
			*c = *m.Data[i][j] * v
			n.Data[i][j] = c
		}
	}
	return n
}

func (m *immutable) Add(o M) M {
	{
		oR, oC := o.Dim()
		if m.R != oR || m.C != oC {
			panic("dimesion mismatch, can only add matricies of the same dimentions")
		}
	}
	n := new(immutable)
	n.R = m.R
	n.C = m.C
	n.Data = make([][]*complex128, n.R)
	for i := range n.Data {
		n.Data[i] = make([]*complex128, n.C)
		for j := range n.Data[i] {
			c := new(complex128)
			*c = *m.Data[i][j] + o.Get(i, j)
			n.Data[i][j] = c
		}
	}
	return n
}

func (m *immutable) Transpose() M {
	return transpose{
		wrap: m.copy(),
	}
}

func (m *immutable) Dot(B M) M {
	return dot(m, B, m)
}

func (m *immutable) Build(d [][]complex128) M {
	return NewImmutable(d)
}
