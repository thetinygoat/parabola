package main

import (
	"log"
)

func main() {

	srv, err := InitServer("tcp", ":8080", 5)
	if err != nil {
		log.Fatal(err)
	}

	err = srv.Listen()
	if err != nil {
		log.Fatal(err)
	}
}
