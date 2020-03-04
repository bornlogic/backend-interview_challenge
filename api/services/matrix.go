package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/marciusvinicius/Interview-Backend-Code-Challenge/pkg/matrix"
)

type params struct {
	Rows [][]int `json:"rows"`
}

//WhereIsTriangular TODO: ADd comment for
func WhereIsTriangular(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var p params
	err := decoder.Decode(&p)

	if err != nil {
		panic(err)
	}

	fmt.Println(p.Rows)

	builder := matrix.Builder{}

	for r := range p.Rows {
		builder = append(make(matrix.Builder, r))
	}

	m, bErr := matrix.Build(builder)

	if bErr != nil {
		json.NewEncoder(w).Encode(err)
	}

	println(m.String())

	json.NewEncoder(w).Encode(m.TypeOfMatrix())
}
