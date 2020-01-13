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

