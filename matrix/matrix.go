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

// IsDiagonal checks if is a diagonal matrix
func IsDiagonal(m Matrix) bool {
	if !IsSquare(m) {
		return false
	}
	for i, line := range m {
		for j, v := range line {
			if i != j && v != 0 {
				return false
			}
		}
	}
	return true
}






