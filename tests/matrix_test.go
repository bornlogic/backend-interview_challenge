package testing

import (
	"testing"

	"github.com/marciusvinicius/InterView-Backend-Code/pkg/matrix"
)

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
	println(r)
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
	println(r)
	if r == true {
		t.Errorf("Matrix shouldn`t be Squared")
	}
}
