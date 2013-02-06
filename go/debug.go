package main

import "fmt"

func debug(s ...interface{}) {
	fmt.Println(s...)
}

func main() {
	//look ma, no fmt
	debug(42, "yeah", []string{"a", "b"},
		map[int]int{1: 1, 2: 2})
}
