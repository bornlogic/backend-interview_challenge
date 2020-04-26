package handlers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/zbioe/ibcc/api/handlers"
)

type testMatrixTest struct {
	req  *http.Request
	code int
	ct   string
}

var testMatrixTests = []testMatrixTest{
	// [0] returns 200 to valid triangular matrix
	testMatrixTest{
		req: httptest.NewRequest(
			http.MethodPost, "/testMatrix",
			strings.NewReader(
				`{"matrix": [[1,1,0],
                                             [0,1,1],
                                             [0,0,0]], "testName": "triangular"}`),
		),
		code: http.StatusOK,
		ct:   "application/json",
	},

	// [1] returns bad request to nil body sended
	testMatrixTest{
		req:  httptest.NewRequest(http.MethodPost, "/testMatrix", nil),
		code: http.StatusBadRequest,
		ct:   "application/json",
	},

	// [2] returns bad request for invalid json
	testMatrixTest{
		req: httptest.NewRequest(
			http.MethodPost, "/testMatrix",
			strings.NewReader(
				`{`),
		),
		code: http.StatusBadRequest,
		ct:   "application/json",
	},

	// [3] returns bad request for missing matrix
	testMatrixTest{
		req: httptest.NewRequest(
			http.MethodPost, "/testMatrix",
			strings.NewReader(
				`{}`),
		),
		code: http.StatusBadRequest,
		ct:   "application/json",
	},

	// [4] returns bad request for missing test name
	testMatrixTest{
		req: httptest.NewRequest(
			http.MethodPost, "/testMatrix",
			strings.NewReader(
				`{"matrix": [[0]]}`),
		),
		code: http.StatusBadRequest,
		ct:   "application/json",
	},

	// [5] returns not acceptable for invalid test name
	testMatrixTest{
		req: httptest.NewRequest(
			http.MethodPost, "/testMatrix",
			strings.NewReader(
				`{"matrix": [[0]], "testName": "invalid"}`),
		),
		code: http.StatusNotAcceptable,
		ct:   "application/json",
	},

	// [6] returns not acceptable for invalid triangular matrix
	testMatrixTest{
		req: httptest.NewRequest(
			http.MethodPost, "/testMatrix",
			strings.NewReader(
				`{"matrix": [[1,1]], "testName": "triangular"}`),
		),
		code: http.StatusNotAcceptable,
		ct:   "application/json",
	},
}

func runTestMatrixTest(t testMatrixTest) error {
	testMatrix := handlers.NewTestMatrix()
	rec := httptest.NewRecorder()
	testMatrix.Serve(rec, t.req)
	if rec.Code != t.code {
		return fmt.Errorf("got [%#v], want [%#v]", rec.Code, t.code)
	}
	if ct := rec.HeaderMap["Content-Type"][0]; ct != t.ct {
		return fmt.Errorf("got [%#v], want [%#v]", ct, t.ct)
	}
	return nil
}

func TestTestMatrix(t *testing.T) {
	for i, test := range testMatrixTests {
		if err := runTestMatrixTest(test); err != nil {
			t.Fatalf("[%d]: %s", i, err)
		}
	}
}
