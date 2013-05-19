package main

import (
	"fmt"
)

func main() {
	var a string
	var b [1]int
	var c *string

	fmt.Println("Address of variable a:", &a)
	fmt.Println("Address of array b:", &b)
	fmt.Println("Address of pointer c:", &c)

	d := [1]int{1}

	takeArray(&d)
}

func takeArray(a *[1]int) {
	fmt.Printf("takeArray: Got array: %v at memory address: %p", a, &a)
}
