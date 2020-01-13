package matrix

type (
	Matrix [][]int
)

// IsSquare checks if is a square matrix
func IsSquare(m Matrix) bool {
	nLine := len(m)
	for _, line := range m {
		nColumn := len(line)
		if nColumn != nLine {
			return false
		}
	}
	return true
}



