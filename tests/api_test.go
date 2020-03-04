package testing

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	services "github.com/marciusvinicius/Interview-Backend-Code-Challenge/api/services"
)

func TestWhereIsTriangularUpper(t *testing.T) {
	rows := [][]int{
		[]int{1, 4, 1},
		[]int{0, 6, 4},
		[]int{0, 0, 1},
	}
	data := map[string][][]int{"rows": rows}
	jsonData, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", "/whereistriangular", bytes.NewBuffer(jsonData))

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(services.WhereIsTriangular)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `Upper Triangular`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestWhereIsTriangularLower(t *testing.T) {
	rows := [][]int{
		[]int{1, 0, 0},
		[]int{2, 8, 0},
		[]int{4, 9, 7},
	}

	data := map[string][][]int{"rows": rows}
	jsonData, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", "/whereistriangular", bytes.NewBuffer(jsonData))

	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(services.WhereIsTriangular)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `Lower Triangular`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
