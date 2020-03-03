package main

import (
	"encoding/json"
	"net/http"
)

func getMembers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	members := db.getMembers()
	json.NewEncoder(w).Encode(members)
}
