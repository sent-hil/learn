package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	log.Println("Starting server on :3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Println("Starting server failed", err)
	}
}

func homeHandler(resp http.ResponseWriter, req *http.Request) {
	log.Println("Got request on /")
}
