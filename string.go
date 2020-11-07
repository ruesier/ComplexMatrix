package complexMatrix

import (
	"fmt"
	"strings"
)

func Sprint(m M) string {
	return SprintCustom(m, "\n\t[", "], ", ", ")
}

func SprintCustom(m M, newRow string, endRow string, colSpace string) string {
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
