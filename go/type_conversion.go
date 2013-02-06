package main

import (
	"fmt"
)

func main() {
	var foo interface{}
	foo = "Hello, world!"

	//force foo to be an int, fails
	//fmt.Println(foo.(int))

	//return value of type cast, bool status of type cast
	//val = 0, ok = false
	val, ok := foo.(int)
	fmt.Println(val, ok)

	//val = "Hello World", ok = true
	right_val, ok := foo.(string)
	fmt.Println(right_val, ok)
}
