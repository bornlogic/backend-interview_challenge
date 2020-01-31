package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// dataFromReq read request body and returns content in bytes
func dataFromReq(r *http.Request) ([]byte, error) {
	if r.Body == nil {
		return nil, fmt.Errorf("body is nil")
	}
	defer r.Body.Close()

	return ioutil.ReadAll(r.Body)
}

// badRequest send message bad request with json compatibility
func badRequest(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintln(w, strconv.Quote(msg))
}

func notAcceptable(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusNotAcceptable)
	fmt.Fprintln(w, strconv.Quote(msg))
}

func setupJSONHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
