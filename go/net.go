package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//make http request
	r, err := http.Get("http://google.com/robots.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	r.Body.Close()

	if err == nil {
		fmt.Println(string(b))
	}
}
