package complexMatrix

import (
	"fmt"
	"strings"
)

// Generates multiline string representation of matrix
func SPrintLines(m M) string {
	return SPrintCustom(m, "[", "],\n", ", ")
}

// Generates string representation of matrix. Uses provided element separators.
func SPrintCustom(m M, newRow string, endRow string, colSpace string) string {
	R, C := m.Dim()
	sb := new(strings.Builder)
	for i := 0; i < R; i++ {
		sb.WriteString(newRow)
		for j := 0; j < C-1; j++ {
			sb.WriteString(fmt.Sprint(m.Get(i, j)))
			sb.WriteString(colSpace)
		}
		if C > 0 {
			sb.WriteString(fmt.Sprint(m.Get(i, C-1)))
		}
		sb.WriteString(endRow)
	}
	return sb.String()
}
