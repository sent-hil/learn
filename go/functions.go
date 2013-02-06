package main

import (
	"fmt"
)

type Callback func()

func main() {
	var my_func Callback
	my_func = func() {
		fmt.Println(11)
	}

	my_func()

	funcArray := make([]Callback, 0)
	for i := 0; i < 10; i++ {
		aNewFunc := func() { fmt.Println(i) }
		funcArray = append(funcArray, aNewFunc)
	}

	for i := 0; i < 10; i++ {
		funcArray[i]()
	}

	funcArray = make([]Callback, 0)
	for i := 0; i < 10; i++ {
		aNewFunc := mkCallback(i)
		funcArray = append(funcArray, aNewFunc)
	}

	for i := 0; i < 10; i++ {
		funcArray[i]()
	}
}
func mkCallback(i int) Callback {
	return func() { fmt.Println(i) }
}
