package main

import (
	"log"
	"reflect"
)

type Hello struct {
	Name string
}

func main() {
	hello := Hello{Name: "first"}
	h := reflect.ValueOf(&hello).Elem()

	log.Println(h.FieldByName("Name").String())

	h.FieldByName("Name").SetString("second")

	log.Println(h.FieldByName("Name").String())
}
