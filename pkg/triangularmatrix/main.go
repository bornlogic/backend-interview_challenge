package triangularmatrix

import "fmt"

type Matrix []int
type Row []int
type Builder []Row

var DEBUG bool = false

/*
 * Provide `true` if you want errors to panic
 */
func SetDebug(debug bool) {
	DEBUG = debug
}

func generateError(message string) (err error) {
	if DEBUG {
		panic(message)
	} else {
		err = fmt.Errorf(message)
	}

	return
}

// Rows TODO: Implement tests
func (m Matrix) Rows() int {
	return int(m[0])
}

// Cols TODO: Implement tests
func (m Matrix) Cols() int {
	return int(m[1])
}

func (m Matrix) String() string {
	output := "\n"
	for i := 0; i < int(m[0]); i++ {
		output += "{\t\t"
		for j := 0; j < int(m[1]); j++ {
			output = fmt.Sprintf("%v%v", output, m.At(i, j))
			if j < int(m[1])-1 {
				output += "\t\t"
			}
		}
		output += "\t\t}\n"
	}

	return output
}

// At TODO: Implement tests
func (m Matrix) At(row, col int) int {
	return m[m.IndexFor(row, col)]
}

// IndexFor TODO: Implement tests
func (m Matrix) IndexFor(row, col int) int {
	return row*int(m[1]) + col + 2
}

// GetRow TODO: Implement tests
func (m Matrix) GetRow(index int) (row []int, err error) {
	if index+1 > m.Rows() {
		err = fmt.Errorf("Row %d is out of matrix", row)
		return
	}

	for i := 0; i < m.Cols(); i++ {
		row = append(row, m.At(index, i))
	}

	return
}

// IsSquare TODO: Implement tests
func (m Matrix) isSquare() bool {
	return false
}

// TypeOfMatrix TODO: Implement tests
func (m Matrix) TypeOfMatrix() string {
	if !m.isSquare() {
		return "Matrix isen`t a valid Square Matrix"
	}

	if m.isUpperTriangular() {
		return "Upper Triangular"
	}
	if m.isLowerTriangular() {
		return "Lower Triangular"
	}
	return "Isen`t Triangular"
}

//
func (m Matrix) isUpperTriangular() bool {
	if !m.isSquare() {
		return false
	}

	for i := 0; i < int(m[0]); i++ {
		for j := i + 1; j < int(m[i]); j++ {
			if m[i][j] != 0 {
				return false
			}
		}
	}

	return true
}

//
func (m Matrix) isLowerTriangular() bool {
	if m.isSquare() {
		return false
	}

	for i := 0; i < int(m[0]); i++ {
		for j := i + 1; j < int(m[i]); j++ {
			if m[i][j] != 0 {
				return false
			}
		}
	}

	return true
}
