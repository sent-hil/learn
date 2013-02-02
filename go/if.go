package main

import (
	"fmt"
)

func main() {
	var b bool

	if true {
		b = true
	} else if false {
		b = false
	} else {
		b = false
	}

	fmt.Println(b)
}
