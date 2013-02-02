package main

import "fmt"

type TZ int

const UTC TZ = 0 * 60 * 60
const EST TZ = -5 * 60 * 60

const (
	a int = 1

	//implicitly same as line above, i.e. b = 0
	b
	c
)

const (
	bit0 uint32 = iota

	//iota increases each time it is called
	bit1
	bit2
	bit3
	bit4
)

func main() {
	//print constant
	fmt.Println(UTC, EST)

	//print vars with defaults
	fmt.Println(a, b, c)

	//print self increasing vars
	fmt.Println(bit0, bit1, bit2, bit3, bit4)
}
