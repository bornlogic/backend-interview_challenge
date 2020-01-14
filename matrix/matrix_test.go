package matrix_test

import (
	"testing"
	"github.com/iuryfukuda/ibcc/matrix"
)


// Cases is the abstraction of multiple checks if matrix is
type Cases []struct{
	name string
	m  matrix.Matrix
}

// testCases test a check function over Cases and check if returns the expected boolean
func testCases(t *testing.T, cs Cases, f func(m matrix.Matrix) bool, expectedOk bool) {
	for _, tt := range cs {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ok := f(tt.m)
			if expectedOk != ok {
				t.Fatalf("Got: %t, Want: %t", ok, expectedOk)
			}
		})
	}
}

var testInvalid = Cases{
	{
		"empty",
		matrix.Matrix{},
	},
	{
		"nil",
		nil,
	},

}

var testIsSquare = Cases{
	{
		"square 2x2",
		matrix.Matrix{
			{1, 2},
			{1, 2},
		},
	}, {
		"square 3x3",
		matrix.Matrix{
			{1, 2, 3},
			{1, 2, 3},
			{1, 2, 3},
		},
	},
}

var testNotIsSquare = Cases{
	{
		"missing line",
		matrix.Matrix{
			{1, 2, 3},
			{1, 2, 3},
		},
	},
	{
		"missing column",
		matrix.Matrix{
			{1, 2},
			{1, 2},
			{1, 2},
		},
	},
}

func TestIsSquare(t *testing.T) {
	testCases(t, testInvalid, matrix.IsSquare, false)

	testCases(t, testIsSquare, matrix.IsSquare, true)
	testCases(t, testNotIsSquare, matrix.IsSquare, false)
}

var testNotIsDiagonal = Cases{
	{
		"one more element in upper",
		matrix.Matrix{
			{1, 0, 0},
			{0, 8, 1},
			{0, 0, 7},
		},
	},
	{
		"one more element in lower",
		matrix.Matrix{
			{1, 0, 0},
			{0, 8, 0},
			{0, 1, 7},
		},
	},
}

var testIsDiagonal = Cases{
	{
		"diagonal only numbers",
		matrix.Matrix{
			{1, 0, 0},
			{0, 8, 0},
			{0, 0, 7},
		},
	},
	{
		"diagonal with zeros",
		matrix.Matrix{
			{0, 0, 0},
			{0, 0, 0},
			{0, 0, 1},
		},
	},
}

func TestIsDiagonal(t *testing.T) {
	testCases(t, testInvalid, matrix.IsDiagonal, false)
	testCases(t, testNotIsSquare, matrix.IsDiagonal, false)

	testCases(t, testIsDiagonal, matrix.IsDiagonal, true)
	testCases(t, testNotIsDiagonal, matrix.IsDiagonal, false)
}

var testNotIsUpperTriangular = Cases{
	{
		"one more element in lower",
		matrix.Matrix{
			{1, 1, 0},
			{0, 8, 0},
			{0, 1, 7},
		},
	},
	{
		"two more elements in lower",
		matrix.Matrix{
			{1, 1, 0},
			{1, 8, 3},
			{0, 1, 7},
		},
	},
}

var testIsUpperTriangular = Cases{
	{
		"upper only numbers",
		matrix.Matrix{
			{1, 4, 1},
			{0, 6, 4},
			{0, 0, 1},
		},
	},
	{
		"upper with zeros",
		matrix.Matrix{
			{1, 0, 0},
			{0, 8, 1},
			{0, 0, 0},
		},
	},
}

func TestIsUpperTriangular(t *testing.T) {
	testCases(t, testInvalid, matrix.IsUpperTriangular, false)
	testCases(t, testNotIsSquare, matrix.IsUpperTriangular, false)

	testCases(t, testIsUpperTriangular, matrix.IsUpperTriangular, true)
	testCases(t, testNotIsUpperTriangular, matrix.IsUpperTriangular, false)
}

var testNotIsLowerTriangular = Cases{
	{
		"one more element in upper",
		matrix.Matrix{
			{1, 1, 0},
			{0, 8, 0},
			{0, 1, 7},
		},
	},
	{
		"two more elements in upper",
		matrix.Matrix{
			{1, 1, 0},
			{1, 8, 3},
			{0, 1, 7},
		},
	},
}

var testIsLowerTriangular = Cases{
	{
		"lower only numbers",
		matrix.Matrix{
			{1, 0, 0},
			{2, 8, 0},
			{4, 9, 7},
		},
	},
	{
		"lower with zeros",
		matrix.Matrix{
			{1, 0, 0},
			{0, 3, 0},
			{0, 1, 0},
		},
	},
}

func TestIsLowerTriangular(t *testing.T) {
	testCases(t, testInvalid, matrix.IsLowerTriangular, false)
	testCases(t, testNotIsSquare, matrix.IsLowerTriangular, false)

	testCases(t, testIsLowerTriangular, matrix.IsLowerTriangular, true)
	testCases(t, testNotIsLowerTriangular, matrix.IsLowerTriangular, false)
}

var testNotIsTriangular = Cases{
	{
		"full numbers",
		matrix.Matrix{
			{1, 2, 3},
			{1, 2, 3},
			{1, 2, 3},
		},
	},
	{
		"one more in upper and lower",
		matrix.Matrix{
			{1, 0, 3},
			{1, 2, 0},
			{0, 0, 0},
		},
	},

}

func TestIsTriangular(t *testing.T) {
	testCases(t, testInvalid, matrix.IsTriangular, false)
	testCases(t, testNotIsSquare, matrix.IsTriangular, false)

	testCases(t, testIsDiagonal, matrix.IsTriangular, true)
	testCases(t, testIsUpperTriangular, matrix.IsTriangular, true)
	testCases(t, testIsLowerTriangular, matrix.IsTriangular, true)

	testCases(t, testNotIsTriangular, matrix.IsTriangular, false)
}

func BenchmarkIsTriangular(b *testing.B) {
	benchmarkCases := []Cases{
		testInvalid,
		testNotIsSquare,
		testIsDiagonal,
		testIsUpperTriangular,
		testIsLowerTriangular,
		testNotIsTriangular,
	}
	for _, cases := range benchmarkCases {
		for _, bm := range cases {
			b.Run(bm.name, func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					matrix.IsTriangular(bm.m)
				}
			})
		}
	}
}
