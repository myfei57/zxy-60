package main

import (
	"flag"
	"log"

	"bagsort/internal/console"
)

func main() {
	addr := flag.String("addr", ":8080", "listen address")
	data := flag.String("data", ".", "data directory")
	flag.Parse()

	server, err := console.New(*data)
	if err != nil {
		log.Fatal(err)
	}
	if err := server.Run(*addr); err != nil {
		log.Fatal(err)
	}
}
