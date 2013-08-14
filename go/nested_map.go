package main

import "log"

func main() {
	a := make(map[string]*map[string]string)

	b := a["a"]

	*b = make(*map[string]string)

	log.Println(b)
}
