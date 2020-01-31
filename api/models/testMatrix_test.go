package models_test

import (
	"testing"

	"github.com/iuryfukuda/ibcc/api/models"
	"github.com/iuryfukuda/ibcc/matrix"
)

type testMatrixTest struct {
	tm  models.TestMatrix
	err error
}

var testMatrixTests = []testMatrixTest{{
	models.TestMatrix{
		matrix.Matrix{
			[]int{1, 2, 3},
			[]int{0, 0, 3},
			[]int{0, 0, 3},
		}, "triangular",
	}, nil,
}, {
	models.TestMatrix{
		matrix.Matrix{
			[]int{1, 2},
			[]int{1, 1},
		}, "triangular",
	}, models.ErrorIsNotTriangular,
}, {
	models.TestMatrix{
		matrix.Matrix{
			[]int{1},
		}, "",
	}, models.ErrorMissingTestName,
}, {
	models.TestMatrix{
		matrix.Matrix{
			[]int{1},
		}, "invalid",
	}, models.ErrorInvalidTestName,
}}

func TestTestMatrix(t *testing.T) {
	for i, tt := range testMatrixTests {
		if err := tt.tm.Test(); err != tt.err {
			t.Fatalf("[%d] unexpected error: got [%s] want [%s]", i, err, tt.err)
		}
	}

}
