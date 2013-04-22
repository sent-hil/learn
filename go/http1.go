package main

import (
	"fmt"
	"net/http"
)

func home(a int) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request)
	}
}

func main() {
	var a = 1
	http.HandleFunc("/", home(a))
	http.HandleFunc("/host", host(a))
	http.ListenAndServe(":8000", nil)
}
