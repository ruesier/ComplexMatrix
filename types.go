/*
complexMatrix is a golang package the implements Matrix types with elements of complex numbers

Copyright 2022 Devon Call
*/
package complexMatrix

import "fmt"

// M represents a matrix of complex numbers. It comes in immutable and mutable varieties.
type M interface {

	// Returns the dimensions of the matrix, height then width.
	Dim() (rows int, columns int)

	// Get the value at a particular row and column.
	Get(row int, column int) complex128

	/* Sets the value at a particular row and column, returns the updated matrix.
	Immutable matricies return a new matrix instance with the updated value.
	Mutable matricies return themselves after updating the value.
	*/
	Set(value complex128, row int, column int) M

	// Updates each value in the matrix by multiplying by the parameter.
	// Immutable matricies return a new matrix instance with updated values.
	// Mutable matricies return themselves after updating the values.
	Scale(c complex128) M

	// Returns matrix by adding elements together of each matrix. Both matricies must have the same dimensions.
	// For A.Add(B):
	// If A is immutable it returns a new matrix.
	// If A is mutable, A is updated and A is returned
	Add(matrix M) M

	// Returns a transposed matrix.
	// Immutable matricies return a matrix that is also immutable.
	// Mutable matricies are used as the underlying matrix, meaning that modifications will alter the original matrix.
	// for example, A.Transpose().Set(0, 1, 2) will have the same effect for A.Set(0, 2, 1)
	Transpose() M

	// Matrix Dot Multiplication
	// A.Dot(B) => AB, B.Dot(A) => BA
	// Both immutable and mutable matricies produce new matricies, because the dimensions are not garunteed to stay constant.
	Dot(matrix M) M

	// Returns an updated matrix with each element being the result of calling the passed in function to each original element.
	// Mutable matricies will also be updated.
	Map(f func(v complex128, r int, c int) complex128) M

	// Returns a new matrix with the provided dimensions.
	// If i, j is within the dimensions of the original matrix, then the new matrix will have the same values at those coordinates.
	Resize(Rows int, Columns int) M

	// Returns an immutable version of the matrix
	Immutable() M

	// Returns an mutable version of the matrix
	Mutable() M
}

// Real returns a 2-d array of the real parts of a matrix
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

// Imag returns a 2-d array of the imaginary parts of a matrix
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

// Two matricies are equal if they have the same dimensions and posistion has the same values
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

// Constructor for an immutable matrix
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

// Combines real and imaginary 2-d arrays into an immutable matrix
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
		panic(fmt.Errorf("complexMatrix.immutable.Set parameter error: i (%d) or j(%d) is out of bounds (%d, %d)", i, j, R, C))
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
