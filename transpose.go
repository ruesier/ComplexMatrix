package complexMatrix

type transpose struct {
	wrap M
}

func (t transpose) Dim() (r int, c int) {
	c, r = t.wrap.Dim()
	return
}

func (t transpose) Get(r int, c int) complex128 {
	return t.wrap.Get(c, r)
}

func (t transpose) Set(v complex128, r int, c int) M {
	return t.wrap.Set(v, c, r).Transpose()
}

func (t transpose) Scale(c complex128) M {
	return t.wrap.Scale(c).Transpose()
}

func (t transpose) Add(m M) M {
	return t.wrap.Add(m.Transpose()).Transpose()
}

func (t transpose) Transpose() M {
	switch v := t.wrap.(type) {
	case *immutable:
		return v.copy()
	default:
		return v
	}
}

// A.Dot(B) => AB, B.Dot(A) => BA
func (t transpose) Dot(m M) M {
	switch v := t.wrap.(type) {
	case Builder:
		return dot(t, m, v)
	default:
		return dot(t, m, &immutable{})
	}
}
