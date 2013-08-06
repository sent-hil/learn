package main

import (
	"log"
)

type Person struct {
	name     string `json:"name"`
	fullName map[string]map[string]string
}

func main() {
	firstName := map[string]string{"first": "senthil"}
	fullName := map[string]map[string]string{"first": firstName}
	person := Person{"s", fullName}
	person.fullName["first.first"] = map[string]string{"hello": "hola"}
	log.Println(person.fullName["first.first"])
	log.Println(person.fullName["first"])
}
