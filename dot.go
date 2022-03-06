package complexMatrix

// Builder represents objects that can construct complexMatrix.M.
// Many types that implement M also implement Builder, so that a new instance
// of an M can be created using the same underlying type without cumbersome
// type checking
type builder interface {
	Build([][]complex128) M
}

// dot implementation of matrix dot product, helper function for Dot methods
// TODO: currently implemented with naive solution, want to update with faster algorithm
func dot(A M, B M, b builder) M {
	targetR, L := A.Dim()
	B_L, targetC := B.Dim()
	if L != B_L {
		panic("incorrect dimensions for matrix dot product, inner dimensions do not match")
	}
	target := make([][]complex128, targetR)
	for i := range target {
		target[i] = make([]complex128, targetC)
		for j := range target[i] {
			total := 0 + 0i
			for k := 0; k < L; k++ {
				total += A.Get(i, k) * B.Get(k, j)
			}
			target[i][j] = total
		}
	}
	return b.Build(target)
}
