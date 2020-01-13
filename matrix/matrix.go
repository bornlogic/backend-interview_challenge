// matrix performs some checks about matrix types
package matrix

type (
	Matrix [][]int
)

// IsTriangular checks if is a triangular matrix
func IsTriangular(m Matrix) bool {
	if !IsSquare(m) {
		return false
	}
	return IsDiagonal(m) || IsUpperTriangular(m) || IsLowerTriangular(m)
}

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

// IsUpperTriangular checks if is a upper triangular matrix
func IsUpperTriangular(m Matrix) bool {
	if !IsSquare(m) {
		return false
	}
	for i, line := range m {
		for j, n := range line {
			if i > j && n != 0 {
				return false
			}
		}
	}
	return true
}

// IsLowerTriangular checks if is a upper triangular matrix
func IsLowerTriangular(m Matrix) bool {
	if !IsSquare(m) {
		return false
	}
	for i, line := range m {
		for j, n := range line {
			if i < j && n != 0 {
				return false
			}
		}
	}
	return true
}
