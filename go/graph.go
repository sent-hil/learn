package main

import (
	"fmt"
)

type Node struct{ Name string }
type MapGraph map[Node]map[Node]bool

func main() {
	left := Node{"left"}
	right := Node{"right"}

	graph := MapGraph{left: map[Node]bool{right: true}}

	fmt.Println(left, right)
	fmt.Println(graph)
}
