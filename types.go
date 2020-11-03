package complexMatrix

type M interface{
  Dim() (r int, c int)
  Get(r int, c int) complex128
  Set(c complex128, r int, c int) M
  Scale(c complex128) M
  Add(m M) M
  Transpose() M
  Dot(m M) M // A.Dot(B) => AB, B.Dot(A) => BA
}

// TODO: loose and strict implementations of M

type immutable struct{
  R int
  C int
  Data [][]*complex128
}

func NewImmutable(table [][]complex128) M {
  n := new(immutable)
  n.R = len(table)
  n.Data = make([][]*complex128, n.R)
  n.C = len(table[0])
  for i := range n.Data{
    n.Data[i] = make([]*complex128, n.C)
    if len(table[i]) != n.C {
      panic("Maxtrix should be rectangular")
    }
    for j := range n.Data[i]{
      n.Data[i][j] = &table[i][j]
    }
  }
  return n
}

func (m *immutable) copy() *immutable {
  n = new(immutable)
  *n = *m
  return n
}

func (m *immutable) Dim() (int, int) {
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

func (m *immutable) Scale(c complex128) M {
  n := new(immutable)
  n.R = m.R
  n.C = m.C
  n.Data = make([][]*complex128, n.R)
  for i := range n.Data {
    n.Data[i] = make([]*complex128, n.C)
    for j := range n.Data[i] {
      c := new(complex128)
      *c = m.Data[i][j] * c
      n.Data[i][j] = c
    }
  }
  return n
}

func (m *immutable) Add(o M) M {
  if m.R != o.R || m.C != o.C {
    panic("dimesion mismatch, can only add matricies of the same dimentions")
  }
  n := new(immutable)
  n.R = m.R
  n.C = m.C
  n.Data = make([][]*complex128, n.R)
  for i := range n.Data {
    n.Data[i] = make([]*complex128, n.C)
    for j := range n.Data[i] {
      c := new(complex128)
      *c = m.Data[i][j] + o.Data[i][j]
      n.Data[i][j] = c
    }
  }
  return n
}

func (m *immutable) Transpose() M {
  n := new(immutable)
  n.R = m.C
  n.C = m.R
  n.Data = make([][]*complex128, n.R)
}
Dot(m M) M // A.Dot(B) => AB, B.Dot(A) => BA
