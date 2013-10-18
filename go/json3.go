package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	Name     string `json:"name"`
	FullName int    `json:"fullName"`
}

func main() {
	data := `{"name":"Hello", "fullName":"Hello"}`

	var person Person
	err := json.Unmarshal([]byte(data), &person)

	log.Println(err, person)
}
