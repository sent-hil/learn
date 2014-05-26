package main

import (
	"log"
	"runtime/debug"
)

type A struct {
	Name string
}

type B struct{}

func (b B) No(n string) {
	log.Println(string(debug.Stack()))
}

func main() {
	var a = A{Name: "indiana jones"}
	B{}.No(a.Name)
}
