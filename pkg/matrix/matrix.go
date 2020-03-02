package matrix

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
func (m Matrix) IsSquared() bool {
	return m.Valid() && m.Cols() == m.Rows()
}

// TypeOfMatrix TODO: Implement tests
func (m Matrix) TypeOfMatrix() string {
	if !m.IsSquared() {
		return "Matrix isen`t a valid Square Matrix"
	}

	if m.isUpperTriangular() {
		return "Upper Triangular"
	}
	if m.IsLowerTriangular() {
		return "Lower Triangular"
	}
	return "Isen`t Triangular"
}

// isUpperTriangular TODO: Implement tests
func (m Matrix) isUpperTriangular() bool {
	if !m.IsSquared() {
		return false
	}

	for i := 1; i < m.Cols(); i++ {
		for j := 0; j < i; j++ {
			if m.At(i, j) != 0 {
				return false
			}
		}
	}

	return true
}

// IsLowerTriangular TODO: Implement tests
func (m Matrix) IsLowerTriangular() bool {
	if m.IsSquared() {
		return false
	}

	for i := 0; i < m.Cols(); i++ {
		for j := i + 1; j < m.Rows(); j++ {
			if m.At(i, j) != 0 {
				return false
			}
		}
	}

	return true
}

/*
 * A matrix is valid if it has rows, columns and if all its rows have the same
 * amount of columns.
 */
func (m Matrix) Valid() bool {
	return len(m) > 2 && len(m) == int(m[0])*int(m[1])+2
}

/*
 * Tells if two matrices have the same dimensions and same
 * values in each cell.
 */
func (m Matrix) EqualTo(other Matrix) bool {
	for i := 0; i < len(m); i++ {
		if m[i] != other[i] {
			return false
		}
	}

	return true
}

/*
 *
 * True if both matrices are valid and have the same dimensions
 */
func (m Matrix) SameDimensions(other Matrix) bool {
	if !m.Valid() || !other.Valid() {
		return false
	}

	return m[0] == other[0] && m[1] == other[1]
}

//Build
func Build(builder Builder) (resultMatrix Matrix, err error) {
	if len(builder) == 0 || len(builder[0]) == 0 {
		err = generateError(fmt.Sprintf("Can't build empty matrix. If you want to generate a zero matrix, use GenerateMatrix()"))
		return
	}

	resultMatrix = GenerateMatrix(len(builder), len(builder[0]))
	for i, row := range builder {
		for j, value := range row {
			resultMatrix[resultMatrix.IndexFor(i, j)] = value
		}
	}

	return
}

/*
 * Generate an zero matrix with `rows` rows and `cols` cols
 */
func GenerateMatrix(rows, cols int) (matrix Matrix) {
	matrix = make(Matrix, rows*cols+2)
	matrix[0] = int(rows)
	matrix[1] = int(cols)
	return
}
