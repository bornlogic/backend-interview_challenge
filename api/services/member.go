package api

import (
	"net/http"
)

func GetMembers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//members := db.getMembers()
	//json.NewEncoder(w).Encode(members)
}

func InviteMember(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//members := db.getMembers()
	//json.NewEncoder(w).Encode(members)
}

func Authenticate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//members := db.getMembers()
	//json.NewEncoder(w).Encode(members)
}
