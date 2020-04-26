package main

import (
	"flag"
	"log"

	"github.com/zbioe/ibcc/api"
)

var Port string

func main() {
	flag.Parse()
	server := api.NewServer()
	log.Fatal(server.Run(Port))
}

func init() {
	const (
		usagePort   = "port used for serve"
		defaultPort = ":3000"
	)
	flag.StringVar(&Port, "port", defaultPort, usagePort)
	flag.StringVar(&Port, "p", defaultPort, usagePort+"(shorthand)")
}
