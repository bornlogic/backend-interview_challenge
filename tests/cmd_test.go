package testing

import (
	"testing"

	"github.com/marciusvinicius/Interview-Backend-Code-Challenge/cmd"
)

func TestCallCommand(t *testing.T) {
	//Pass a array of col
	rows := [][]int{
		[]int{1, 2, 3},
		[]int{1, 2, 3},
		[]int{1, 2, 3},
		[]int{1, 2, 3},
	}

	r := cmd.WhereIsTriangular(rows)
	println(r)
	if r != "Isen`t Triangular" {
		t.Errorf("String should be 'Isen`t Triangular'")
	}
}
