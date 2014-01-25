package main

import (
	"log"
)

type aliasToInt int
type secondAlias int

func main() {
	var a aliasToInt
	a = 2

	var b secondAlias
	b = 2

	acceptsInt(1)
	acceptsInt(int(a))
	acceptsInt(int(b))
}

func acceptsInt(al int) {
	log.Println(al)
}
