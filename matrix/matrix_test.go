package matrix_test

import (
	"testing"
	"github.com/iuryfukuda/ibcc/matrix"
)


// Are is the abstraction of multiple checks if matrix is
type Are []struct{
	name string
	m  matrix.Matrix
}

// testAre test a check function over Are and check if returns the expected boolean
func testAre(t *testing.T, a Are, f func(m matrix.Matrix) bool, expectedOk bool) {
	for _, tt := range a {
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

var testInvalid = Are{
	{
		"empty",
		matrix.Matrix{},
	},
	{
		"nil",
		nil,
	},

}

var testIsSquare = Are{
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

var testNotIsSquare = Are{
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
	testAre(t, testInvalid, matrix.IsSquare, false)

	testAre(t, testIsSquare, matrix.IsSquare, true)
	testAre(t, testNotIsSquare, matrix.IsSquare, false)
}

var testNotIsDiagonal = Are{
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

var testIsDiagonal = Are{
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
	testAre(t, testInvalid, matrix.IsDiagonal, false)
	testAre(t, testNotIsSquare, matrix.IsDiagonal, false)

	testAre(t, testIsDiagonal, matrix.IsDiagonal, true)
	testAre(t, testNotIsDiagonal, matrix.IsDiagonal, false)
}

var testNotIsUpperTriangular = Are{
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

var testIsUpperTriangular = Are{
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
	testAre(t, testInvalid, matrix.IsUpperTriangular, false)
	testAre(t, testNotIsSquare, matrix.IsUpperTriangular, false)

	testAre(t, testIsUpperTriangular, matrix.IsUpperTriangular, true)
	testAre(t, testNotIsUpperTriangular, matrix.IsUpperTriangular, false)
}

var testNotIsLowerTriangular = Are{
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

var testIsLowerTriangular = Are{
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
	testAre(t, testInvalid, matrix.IsLowerTriangular, false)
	testAre(t, testNotIsSquare, matrix.IsLowerTriangular, false)

	testAre(t, testIsLowerTriangular, matrix.IsLowerTriangular, true)
	testAre(t, testNotIsLowerTriangular, matrix.IsLowerTriangular, false)
}

var testNotIsTriangular = Are{
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
	testAre(t, testInvalid, matrix.IsTriangular, false)
	testAre(t, testNotIsSquare, matrix.IsTriangular, false)

	testAre(t, testIsDiagonal, matrix.IsTriangular, true)
	testAre(t, testIsUpperTriangular, matrix.IsTriangular, true)
	testAre(t, testIsLowerTriangular, matrix.IsTriangular, true)

	testAre(t, testNotIsTriangular, matrix.IsTriangular, false)
}

func BenchmarkIsTriangular(b *testing.B) {
	benchmarks := []Are{
		testInvalid,
		testNotIsSquare,
		testIsDiagonal,
		testIsUpperTriangular,
		testIsLowerTriangular,
		testNotIsTriangular,
	}
	for _, are := range benchmarks {
		for _, bm := range are {
			b.Run(bm.name, func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					matrix.IsTriangular(bm.m)
				}
			})
		}
	}
}
