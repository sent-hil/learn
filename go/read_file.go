package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("VERSION")
	if err != nil {
		log.Println(err)
	}
	log.Println(strings.TrimSpace(string(file)))

	p := 9001
	data, _ := json.MarshalIndent(p, "", "  ")
	ioutil.WriteFile("VERSION", data, 0644)
}
