package matrix_test

import (
	"testing"
	"github.com/iuryfukuda/ibcc/matrix"
)


// Are is the abstraction of multiple checks if matrix is
type Are []struct{
	name string
	m  matrix.Matrix
	is bool
}

func testAre(t *testing.T, a Are, f func(m matrix.Matrix) bool) {
	for _, tt := range a {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := f(tt.m)
			if is != tt.is {
				t.Fatalf("Got: %t, Want: %t", is, tt.is)
			}
		})
	}
}

var testIsSquare = Are{
	{
		"missing line",
		matrix.Matrix{
			{1, 2, 3},
			{1, 2, 3},
		},
		false,
	},
	{
		"missing column",
		matrix.Matrix{
			{1, 2},
			{1, 2},
			{1, 2},
		},
		false,
	},
	{
		"is",
		matrix.Matrix{
			{1, 2, 3},
			{1, 2, 3},
			{1, 2, 3},
		},
		true,
	},
}

func TestIsSquare(t *testing.T) {
	testAre(t, testIsSquare, matrix.IsSquare)
}

var testIsDiagonal = Are{
	{
		"one more element in upper",
		matrix.Matrix{
			{1, 0, 0},
			{0, 8, 1},
			{0, 0, 7},
		},
		false,
	},
	{
		"one more element in lower",
		matrix.Matrix{
			{1, 0, 0},
			{0, 8, 0},
			{0, 1, 7},
		},
		false,
	},
	{
		"is",
		matrix.Matrix{
			{1, 0, 0},
			{0, 8, 0},
			{0, 0, 7},
		},
		true,
	},
	{
		"is with many zero",
		matrix.Matrix{
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 1},
		},
		true,
	},
}

func TestIsDiagonal(t *testing.T) {
	testAre(t, testIsDiagonal, matrix.IsDiagonal)
}

var testIsUpperTriangular = Are{
	{
		"one more element in lower",
		matrix.Matrix{
			{1, 1, 0},
			{0, 8, 0},
			{0, 1, 7},
		},
		false,
	},
	{
		"two more elements in lower",
		matrix.Matrix{
			{1, 1, 0},
			{1, 8, 3},
			{0, 1, 7},
		},
		false,
	},
	{
		"is",
		matrix.Matrix{
			{1, 4, 1},
			{0, 6, 4},
			{0, 0, 1},
		},
		true,
	},
	{
		"is with many zeroes",
		matrix.Matrix{
			{1, 0, 0},
			{0, 8, 1},
			{0, 0, 0},
		},
		true,
	},
}

func TestIsUpperTriangular(t *testing.T) {
	testAre(t, testIsUpperTriangular, matrix.IsUpperTriangular)
}

var testIsLowerTriangular = Are{
	{
		"one more element in upper",
		matrix.Matrix{
			{1, 1, 0},
			{0, 8, 0},
			{0, 1, 7},
		},
		false,
	},
	{
		"two more elements in upper",
		matrix.Matrix{
			{1, 1, 0},
			{1, 8, 3},
			{0, 1, 7},
		},
		false,
	},
	{
		"is",
		matrix.Matrix{
			{1, 0, 0},
			{2, 8, 0},
			{4, 9, 7},
		},
		true,
	},
	{
		"is with many zeroes",
		matrix.Matrix{
			{1, 0, 0},
			{0, 3, 0},
			{0, 1, 0},
		},
		true,
	},
}

func TestIsLowerTriangular(t *testing.T) {
	testAre(t, testIsLowerTriangular, matrix.IsLowerTriangular)
}



var testIsTriangular = Are{
	{
		"one more element in upper and lower",
		matrix.Matrix{
			{1, 1, 0},
			{0, 8, 0},
			{0, 1, 7},
		},
		false,
	},
	{
		"two more elements in upper and lower",
		matrix.Matrix{
			{1, 1, 0},
			{1, 8, 3},
			{0, 1, 7},
		},
		false,
	},
	{
		"is lower",
		matrix.Matrix{
			{1, 0, 0},
			{2, 8, 0},
			{4, 9, 7},
		},
		true,
	},
	{
		"is upper",
		matrix.Matrix{
			{1, 1, 0},
			{0, 3, 0},
			{0, 0, 0},
		},
		true,
	},
	{
		"is diagonal",
		matrix.Matrix{
			{1, 0, 0},
			{0, 3, 0},
			{0, 0, 0},
		},
		true,
	},
}

func TestIsTriangular(t *testing.T) {
	testAre(t, testIsTriangular, matrix.IsTriangular)
}




