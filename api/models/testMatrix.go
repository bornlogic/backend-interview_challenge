package models

import (
	"fmt"
	"github.com/zbioe/ibcc/matrix"
)

// TestMatrix is the main abstraction about tests over matrix
type TestMatrix struct {
	// matrix is sended matrix in input
	Matrix matrix.Matrix `json:"matrix"`
	// testName refers to type of test used in test of matrix
	TestName string `json:"testName"`
}

var (
	ErrorIsNotTriangular = fmt.Errorf("is not triangular")
	ErrorMissingTestName = fmt.Errorf("missing test name")
	ErrorInvalidTestName = fmt.Errorf("invalid test name")
)

func (tm *TestMatrix) Test() error {
	switch tm.TestName {
	case "triangular":
		if !matrix.IsTriangular(tm.Matrix) {
			return ErrorIsNotTriangular
		}
	// missing testName returns invalid
	case "":
		return ErrorMissingTestName
	// unknow test passed returns invalid
	default:
		return ErrorInvalidTestName
	}
	return nil
}
