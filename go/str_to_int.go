package main

import (
	"fmt"
	"strconv"
)

func main() {
	st1 := strconv.FormatInt(10, 10)
	by1 := []byte(st1)

	//st2 := "A a å Å Â"
	//by2 := []byte(st2)

	st3 := strconv.FormatFloat(10.0, 'b', 3, 64)
	by3 := []byte(st3)

	fmt.Println("Int bytes:", by1)
	//fmt.Println("String bytes:", by2)
	fmt.Println("Float bytes:", by3)

	//fmt.Println("Modulo operator:", 12 % 11)
}
