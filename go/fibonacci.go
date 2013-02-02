package main

import (
	"fmt"
)

func fib(num int) (result int) {
	if num > 2 {
		result = fib(num-1) + fib(num-2)
	} else {
		result = 1
	}

	//returns name of variable in func def, i.e. result
	return
}

func main() {
	for i := 1; i < 20; i++ {
		fmt.Println("Fib:", fib(i))
	}
}
