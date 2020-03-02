package testing

import (
	"matrix"
	"testing"
)

func MatrixIsSquared(t *testing.T) {
	m1, err := matrix.Build(
		matrix.Builder{
			matrix.Row{1, 1, 1},
			matrix.Row{1, 1, 1},
			matrix.Row{1, 1, 1},
		},
	)
	if err != nil {
		println("You passed a 0x0 or 1x0 matrix.")
	}
	r := m1.isSquared()
	r == true
}
