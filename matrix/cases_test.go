package matrix_test

import (
	"github.com/iuryfukuda/ibcc/matrix"
)

// Cases is the abstraction of multiple checks if matrix is
type Cases []struct{
	name string
	m  matrix.Matrix
}

// casesInvalid join all inputs not considered matrix
var casesInvalid = Cases{
	{
		"empty",
		matrix.Matrix{},
	},
	{
		"nil",
		nil,
	},

}

// square matrix cases
var casesIsSquare = Cases{
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

var casesNotIsSquare = Cases{
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

// diagonal matrix cases
var casesIsDiagonal = Cases{
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

var casesNotIsDiagonal = Cases{
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

// upper triangular matrix cases
var casesIsUpperTriangular = Cases{
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

var casesNotIsUpperTriangular = Cases{
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

// lower triangular matrix cases
var casesIsLowerTriangular = Cases{
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

var casesNotIsLowerTriangular = Cases{
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

// triangular matrix cases
var casesNotIsTriangular = Cases{
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
