package main

import (
	"fmt"
	"strconv"
)

func main() {
	st1 := strconv.FormatInt(10, 10)
	by1 := []byte(st1)
	fmt.Println("Int bytes:", st1, by1)

	st2 := "A a å Å Â"
	by2 := []byte(st2)
	fmt.Println("String bytes:", st2, by2)

	st3 := strconv.FormatFloat(10.0, 'b', 3, 64)
	by3 := []byte(st3)
	fmt.Println("Float bytes:", st3, by3)

	fmt.Println("Modulo operator:", 12%11)
}
