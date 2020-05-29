package main

import (
	"log"
	"net/http"
)

// Port - Port number of HTTP server
const Port string = ":8080"

func main() {

	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)

	log.Println("Listening on port", Port)
	err := http.ListenAndServe(Port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
