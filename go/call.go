package main

import (
	"log"
	"runtime/debug"
	"strings"
)

type A struct {
	name string
}

func (a A) Name() string {
	return "Indiana Jones"
}

type B struct{}

func (b B) No(n string) {
	var output = string(debug.Stack())
	var newStr = strings.Split(output, "\n")
	var trimmed = strings.TrimSpace(newStr[3])
	var split = strings.Split(trimmed, ".")
	var splitMore = split[len(split)-1]
	var finally = splitMore[:len(splitMore)-1]

	log.Println(output, "\n", finally)
}

func main() {
	var a = A{}
	B{}.No(a.Name())
}
