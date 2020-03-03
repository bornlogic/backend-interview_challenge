package main

import (
	"os"
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	a := App{}
	a.Initialize(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))
	)
	a.Run(":80")
}
