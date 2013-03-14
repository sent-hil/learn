package main

import (
	"fmt"
	"reflect"
)

type My struct {
	A int `abc:"type14"`
}

func main() {
	x := &My{A: 34}
	f, _ := reflect.TypeOf(x).Elem().FieldByName("A")
	fmt.Println(f.Tag)
}
