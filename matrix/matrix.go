// matrix performs some checks about matrix types
package matrix

import (
	"fmt"
	"strconv"
	"strings"
	"encoding/csv"
)

// Matrix is the main abstraction about type matrix of gemometry
// implements "flag"
type Matrix [][]int

func(m Matrix) String() string {
	var msg string
	for _, line := range m {
		for _, col := range line {
			msg += fmt.Sprintf("%d,", col)
		}
		msg = strings.Trim(msg, ",") + "\n"
	}
	msg = strings.TrimSpace(msg)
	return msg
}

// SetMatrixErrors is the abstraction about errors on set new matrix
type SetMatrixErrors []error
func (es SetMatrixErrors) Error() string {
	var msg string
	for _, e := range es {
		msg = msg + "\n " + e.Error()
	}
	return msg
}

func (es *SetMatrixErrors) add(e error, i, j int) {
	*es = append(*es, fmt.Errorf("(%d, %d): %s", i, j, e))
}

// Set convert csv string passed as string for matrix
func (m *Matrix) Set (s string) error {
	r := csv.NewReader(strings.NewReader(s))
	strMatrix, err := r.ReadAll()
	if err != nil {
		return err
	}

	var errors SetMatrixErrors
	var tmpMatrix = make(Matrix, len(strMatrix))
	for i, line := range strMatrix {
		tmpMatrix[i] = make([]int, len(line))
		for j, col := range line {
			n, err := strconv.Atoi(col)
			if err != nil {
				errors.add(err, i, j)
				continue
			}
			tmpMatrix[i][j] = n
		}
	}

	if len(errors) > 0 {
		return errors
	}

	*m = tmpMatrix
	return nil
}


// IsTriangular checks if is a triangular matrix
func IsTriangular(m Matrix) bool {
	return IsDiagonal(m) || IsUpperTriangular(m) || IsLowerTriangular(m)
}

// IsSquare checks if is a square matrix
func IsSquare(m Matrix) bool {
	// do not accept empty or nil matrix
	if len(m) == 0 {
		return false
	}
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

// IsLowerTriangular checks if is a lower triangular matrix
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
