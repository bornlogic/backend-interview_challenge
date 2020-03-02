package testing

import (
	_ "pkg/triangularmatrix"
	"testing"
)

func MatrixIsSquared(t *testing.T) {
	m1 := [3][3]int{
		[3]int{1, 1, 1},
		[3]int{1, 1, 1},
		[3]int{1, 1, 1},
	}
	r := IsQuare(m1)
	r == true
}

func MatrixIsentSquared(*testing.T) {
	m1 := [3][2]int{
		[3]int{1, 1},
		[3]int{1, 1},
		[3]int{1, 1},
	}
	r := IsSquaredMatrix(m1)
	test.assertFalse(r)
}

func MatrixIsUpperTriangular(*testing.T) {
	m1 := [3][3]int{
		[3]int{1, 1, 1},
		[3]int{1, 1, 1},
		[3]int{1, 1, 1},
	}

	r := WhereIsTriangular(m1)

}

// func MatrixIsLowerriangular(*testing.T) {
// 	m1 := [3][3]int{
// 		[3]int{1, 1, 1},
// 		[3]int{1, 1, 1},
// 		[3]int{1, 1, 1},
// 	}

// 	r := WhereIsTriangular(m1)

// }

// func MatrixIsentTriangular(*testing.T) {
// 	m1 := [3][3]int{
// 		[3]int{1, 1, 1},
// 		[3]int{1, 1, 1},
// 		[3]int{1, 1, 1},
// 	}

// 	r := WhereIsTriangular(m1)
// }

// func MatrixIsInvalid(*testing.T) {
// 	m1 := [3][3]int{
// 		[3]int{1, 1, 1},
// 		[3]int{1, 1, 1},
// 		[3]int{1, 1, 1},
// 	}

// 	r := WhereIsTriangular(m1)

// }
