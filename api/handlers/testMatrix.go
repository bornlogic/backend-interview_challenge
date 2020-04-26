package handlers

import (
	"encoding/json"
	"github.com/zbioe/ibcc/api/models"
	"net/http"
)

type testMatrix struct{}

func NewTestMatrix() *testMatrix {
	return &testMatrix{}
}

func (tm *testMatrix) Serve(w http.ResponseWriter, r *http.Request) {
	setupJSONHeaders(w)
	d, err := dataFromReq(r)
	if err != nil {
		badRequest(w, err.Error())
		return
	}

	var mTM models.TestMatrix
	if err := json.Unmarshal(d, &mTM); err != nil {
		badRequest(w, "invalid body: "+err.Error())
		return
	}

	if len(mTM.Matrix) == 0 {
		badRequest(w, "missing matrix")
		return
	}

	if mTM.TestName == "" {
		badRequest(w, "missing test name")
		return
	}

	if err := mTM.Test(); err != nil {
		notAcceptable(w, "matrix test failed: "+err.Error())
		return
	}
}
