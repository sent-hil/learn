package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x interface{} = 5
	v := reflect.ValueOf(x)

	fmt.Println("val:", v, x)
	fmt.Printf("%T", x)
}
