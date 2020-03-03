package main

import (
	"os"
	"encoding/json"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
}


func main() {
	a := App{}
	a.Initialize(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))
	)
	a.Router.HandleFunc("/", HomeHandler)
	a.Run(":80")
}
