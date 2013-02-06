package main

import (
	"fmt"
)

func hello_there() string {
	return "Hello World"
}

//accepts func that returns string
type str struct {
	fu func() string
}

//recursive type
type StateFunc func(int) StateFunc

func hello(sf StateFunc) {
	fmt.Println(sf)
}

func main() {
	fmt.Println(hello_there())
	var s = str{fu: hello_there}
	fmt.Println(s.fu())

	//create func that returns func which takes string
	B := func(x string) func(string) {
		return func(y string) {
			fmt.Println("a", x, y)
		}
	}

	//call B with "b" and then call result with "c"
	B("b")("c")

	var f StateFunc
	f = func(y int) StateFunc {
		fmt.Println(y)
		return f
	}

	hello(f(1)(2)(3))
}
