package main

import (
	"os"

	"github.com/marciusvinicius/Interview-Backend-Code-Challenge/api"
)

func main() {
	a := api.App{}
	a.Initialize(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	a.Run(":80")
}
