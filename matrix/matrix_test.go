package matrix_test

import (
	"github.com/iuryfukuda/ibcc/matrix"
	"testing"
)

func TestSet(t *testing.T) {
	rawMatrix := `1,1,3
1,3,"3"
1,"1",1`
	var m matrix.Matrix
	err := m.Set(rawMatrix)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSetMatrixErrors(t *testing.T) {
	rawMatrix := `1,1,c
1,3,3
1,c,1`
	var m matrix.Matrix
	err := m.Set(rawMatrix)
	errors, ok := err.(matrix.SetMatrixErrors)
	if !ok {
		t.Fatalf(
			"unexpected type error returned: got: %T, want: %T",
			err, *new(matrix.SetMatrixErrors),
		)
	}

	expectedLen := 2
	if count := len(errors); count != expectedLen {
		t.Fatalf(
			"unexpected len errors: got: %d, want: %d",
			count, expectedLen,
		)
	}
}

// testCases test a check function over Cases and check if returns the expected boolean
func testCases(t *testing.T, cs Cases, f func(m matrix.Matrix) bool, expectedOk bool) {
	for _, tt := range cs {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ok := f(tt.m)
			if expectedOk != ok {
				t.Fatalf("got: %t, want: %t", ok, expectedOk)
			}
		})
	}
}

func TestIsSquare(t *testing.T) {
	testCases(t, casesInvalid, matrix.IsSquare, false)
	testCases(t, casesNotIsSquare, matrix.IsSquare, false)

	testCases(t, casesIsSquare, matrix.IsSquare, true)
}

func TestIsDiagonal(t *testing.T) {
	testCases(t, casesInvalid, matrix.IsDiagonal, false)
	testCases(t, casesNotIsSquare, matrix.IsDiagonal, false)
	testCases(t, casesNotIsDiagonal, matrix.IsDiagonal, false)

	testCases(t, casesIsDiagonal, matrix.IsDiagonal, true)
}

func TestIsUpperTriangular(t *testing.T) {
	testCases(t, casesInvalid, matrix.IsUpperTriangular, false)
	testCases(t, casesNotIsSquare, matrix.IsUpperTriangular, false)
	testCases(t, casesNotIsUpperTriangular, matrix.IsUpperTriangular, false)

	testCases(t, casesIsUpperTriangular, matrix.IsUpperTriangular, true)
}

func TestIsLowerTriangular(t *testing.T) {
	testCases(t, casesInvalid, matrix.IsLowerTriangular, false)
	testCases(t, casesNotIsSquare, matrix.IsLowerTriangular, false)
	testCases(t, casesNotIsLowerTriangular, matrix.IsLowerTriangular, false)

	testCases(t, casesIsLowerTriangular, matrix.IsLowerTriangular, true)
}

func TestIsTriangular(t *testing.T) {
	testCases(t, casesInvalid, matrix.IsTriangular, false)
	testCases(t, casesNotIsSquare, matrix.IsTriangular, false)
	testCases(t, casesNotIsTriangular, matrix.IsTriangular, false)

	testCases(t, casesIsDiagonal, matrix.IsTriangular, true)
	testCases(t, casesIsUpperTriangular, matrix.IsTriangular, true)
	testCases(t, casesIsLowerTriangular, matrix.IsTriangular, true)
}

func BenchmarkIsTriangular(b *testing.B) {
	setOfCases := []Cases{
		casesInvalid,
		casesNotIsSquare,
		casesNotIsTriangular,
		casesIsDiagonal,
		casesIsUpperTriangular,
		casesIsLowerTriangular,
	}
	for _, cases := range setOfCases {
		for _, bm := range cases {
			b.Run(bm.name, func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					matrix.IsTriangular(bm.m)
				}
			})
		}
	}
}
