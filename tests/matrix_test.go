package testing

import (
	"testing"

	"github.com/marciusvinicius/InterView-Backend-Code/pkg/matrix"
)

func TestRows(t *testing.T) {
	m1, err := matrix.Build(
		matrix.Builder{
			matrix.Row{1, 2, 3},
		},
	)

	if err != nil {
		t.Errorf("You passed a 0x0 or 1x0 matrix.")
	}

	r := m1.Rows()
	if r != 1 {
		t.Errorf("Matrix have 1 row")
	}
}

func TestCols(t *testing.T) {
	m1, err := matrix.Build(
		matrix.Builder{
			matrix.Row{1, 2, 3},
		},
	)

	if err != nil {
		t.Errorf("You passed a 0x0 or 1x0 matrix.")
	}

	r := m1.Cols()
	if r != 3 {
		t.Errorf("Matrix have 3 cols")
	}
}

func TestMatrixString(t *testing.T) {
	m1, err := matrix.Build(
		matrix.Builder{
			matrix.Row{1, 2, 3},
		},
	)

	if err != nil {
		t.Errorf("You passed a 0x0 or 1x0 matrix.")
	}

	stm := "\n{\t\t1\t\t2\t\t3\t\t}\n"

	if m1.String() != stm {
		t.Errorf("Shoul be " + stm)
	}
}

func TestMatrixAt(t *testing.T) {
	m1, err := matrix.Build(
		matrix.Builder{
			matrix.Row{1, 2, 3},
			matrix.Row{4, 5, 6},
			matrix.Row{7, 8, 9},
		},
	)
	if err != nil {
		t.Errorf("You passed a 0x0 or 1x0 matrix.")
	}

	r := m1.At(1, 1)
	if r != 5 {
		t.Errorf("Matrix at 1x1 should be 5")
	}
}

func TestMatrixAt2(t *testing.T) {
	m1, err := matrix.Build(
		matrix.Builder{
			matrix.Row{1, 2, 3},
			matrix.Row{4, 5, 6},
			matrix.Row{7, 8, 9},
		},
	)
	if err != nil {
		t.Errorf("You passed a 0x0 or 1x0 matrix.")
	}

	r := m1.At(0, 1)
	if r != 2 {
		t.Errorf("Matrix at 1x1 should be 2")
	}
}

func TestMatrixIndexFor(t *testing.T) {
	m1, err := matrix.Build(
		matrix.Builder{
			matrix.Row{1, 1, 1},
			matrix.Row{1, 1, 1},
			matrix.Row{1, 1, 1},
		},
	)
	if err != nil {
		t.Errorf("You passed a 0x0 or 1x0 matrix.")
	}

	r := m1.IndexFor(0, 1)

	if r != 3 {
		t.Errorf("Matrix index should be 3")
	}
}

func TestMatrixEqualTo(t *testing.T) {

	m1, _ := matrix.Build(
		matrix.Builder{
			matrix.Row{1, 1, 1},
			matrix.Row{1, 1, 1},
			matrix.Row{1, 1, 1},
		},
	)

	m2, _ := matrix.Build(
		matrix.Builder{
			matrix.Row{1, 1, 1},
			matrix.Row{1, 1, 1},
			matrix.Row{1, 1, 1},
		},
	)

	if !m1.EqualTo(m2) {
		t.Errorf("Matrixs should be equal")
	}
}

func TestMatrixDifferentFrom(t *testing.T) {

	m1, _ := matrix.Build(
		matrix.Builder{
			matrix.Row{1, 1, 1},
			matrix.Row{1, 1, 1},
			matrix.Row{1, 1, 1},
		},
	)

	m2, _ := matrix.Build(
		matrix.Builder{
			matrix.Row{1, 2, 1},
			matrix.Row{1, 3, 1},
			matrix.Row{2, 1, 1},
		},
	)

	if m1.EqualTo(m2) {
		t.Errorf("Matrixs should be diferent")
	}
}

func TestMatrixSameDimensions(t *testing.T) {
	m1, _ := matrix.Build(
		matrix.Builder{
			matrix.Row{1, 1, 1},
			matrix.Row{1, 1, 1},
			matrix.Row{1, 1, 1},
		},
	)

	m2, _ := matrix.Build(
		matrix.Builder{
			matrix.Row{1, 2, 1},
			matrix.Row{1, 3, 1},
			matrix.Row{2, 1, 1},
		},
	)

	if !m1.SameDimensions(m2) {
		t.Errorf("Matrixs should have same dimension")
	}
}

func TestMatrixDiferentDimensions(t *testing.T) {
	m1, _ := matrix.Build(
		matrix.Builder{
			matrix.Row{1, 1, 1},
			matrix.Row{1, 1, 1},
		},
	)

	m2, _ := matrix.Build(
		matrix.Builder{
			matrix.Row{1, 2, 1},
			matrix.Row{1, 3, 1},
			matrix.Row{2, 1, 1},
		},
	)

	if m1.SameDimensions(m2) {
		t.Errorf("Matrixs shouldn`t have same dimension")
	}
}

func TestMatrixIsSquared(t *testing.T) {
	m1, err := matrix.Build(
		matrix.Builder{
			matrix.Row{1, 1, 1},
			matrix.Row{1, 1, 1},
			matrix.Row{1, 1, 1},
		},
	)
	if err != nil {
		t.Errorf("You passed a 0x0 or 1x0 matrix.")
	}
	r := m1.IsSquared()

	if r != true {
		t.Errorf("Matrix should be Squared")
	}
}
func TestMatrixIsentSquared(t *testing.T) {
	m1, err := matrix.Build(
		matrix.Builder{
			matrix.Row{1, 1},
			matrix.Row{1, 1},
			matrix.Row{1, 1},
		},
	)
	if err != nil {
		t.Errorf("You passed a 0x0 or 1x0 matrix.")
	}
	r := m1.IsSquared()

	if r == true {
		t.Errorf("Matrix shouldn`t be Squared")
	}
}

func TestIsUpperTriangular(t *testing.T) {
	/** Ex Upper Triangular Matrix
		1 2 3
		0 5 6
		0 0 9
	**/
	m1, err := matrix.Build(
		matrix.Builder{
			matrix.Row{1, 2, 3},
			matrix.Row{0, 5, 6},
			matrix.Row{0, 0, 9},
		},
	)
	if err != nil {
		t.Errorf("You passed a 0x0 or 1x0 matrix.")
	}
	r := m1.IsUpperTriangular()
	if !r {
		t.Errorf("Matrix should be upper triangular")
	}
}

func TestIsUpperTriangular2(t *testing.T) {
	/** Ex Upper Triangular Matrix
		1 4 1
		0 6 4
		0 0 1
	**/
	m1, err := matrix.Build(
		matrix.Builder{
			matrix.Row{1, 4, 1},
			matrix.Row{0, 6, 4},
			matrix.Row{0, 0, 1},
		},
	)
	if err != nil {
		t.Errorf("You passed a 0x0 or 1x0 matrix.")
	}
	r := m1.IsUpperTriangular()
	if !r {
		t.Errorf("Matrix should be upper triangular")
	}
}
func TestIsLowerTriangular(t *testing.T) {
	/** Ex Lower Triangular Matrix
		1 0 0
		4 5 0
		7 8 9
	**/
	m1, err := matrix.Build(
		matrix.Builder{
			matrix.Row{1, 0, 0},
			matrix.Row{4, 5, 0},
			matrix.Row{7, 8, 9},
		},
	)
	if err != nil {
		t.Errorf("You passed a 0x0 or 1x0 matrix.")
	}

	r := m1.IsLowerTriangular()
	if !r {
		t.Errorf("Matrix should be lower triangular")
	}
}
func TestIsLowerTriangular2(t *testing.T) {
	/** Ex Lower Triangular Matrix
		1 0 0
		2 8 0
		4 9 7
	**/
	m1, err := matrix.Build(
		matrix.Builder{
			matrix.Row{1, 0, 0},
			matrix.Row{2, 8, 0},
			matrix.Row{4, 9, 7},
		},
	)
	if err != nil {
		t.Errorf("You passed a 0x0 or 1x0 matrix.")
	}

	r := m1.IsLowerTriangular()

	if !r {
		t.Errorf("Matrix should be lower triangular")
	}
}
func TestTypeOfMatrixLower(t *testing.T) {
	m1, _ := matrix.Build(
		matrix.Builder{
			matrix.Row{1, 0, 0},
			matrix.Row{2, 8, 0},
			matrix.Row{4, 9, 7},
		},
	)

	if m1.TypeOfMatrix() != "Lower Triangular" {
		t.Errorf("String shoubd be 'Lower Triangular'")
	}
}

func TestTypeOfMatrixUpper(t *testing.T) {
	m1, _ := matrix.Build(
		matrix.Builder{
			matrix.Row{1, 4, 1},
			matrix.Row{0, 6, 4},
			matrix.Row{0, 0, 1},
		},
	)

	if m1.TypeOfMatrix() != "Upper Triangular" {
		t.Errorf("String shoubd be 'Upper Triangular'")
	}
}

func TestTypeOfMatrixNoTriangular(t *testing.T) {
	m1, _ := matrix.Build(
		matrix.Builder{
			matrix.Row{1, 4, 1},
			matrix.Row{1, 6, 4},
			matrix.Row{0, 1, 1},
		},
	)

	if m1.TypeOfMatrix() != "Isen`t Triangular" {
		t.Errorf("String shoubd be 'Isen`t Triangular'")
	}
}

func TestTypeOfMatrixNoSquared(t *testing.T) {
	m1, _ := matrix.Build(
		matrix.Builder{
			matrix.Row{1, 4, 1},
			matrix.Row{1, 6, 4},
		},
	)

	if m1.TypeOfMatrix() != "Matrix isen`t a valid Squared Matrix" {
		t.Errorf("String shoubd be 'Matrix isen`t a valid Squared Matrix'")
	}
}
