package main

import "fmt"

type List interface {
	Len() int
}
type LinkedList struct{}
type DoubleLinkedList struct{}

func (ll LinkedList) Len() int       { return 1 }
func (dl DoubleLinkedList) Len() int { return 2 }

func lengthOfList(l List) int { return l.Len() }

func main() {
	var ll LinkedList
	var dl DoubleLinkedList
	fmt.Println(lengthOfList(ll))
	fmt.Println(lengthOfList(dl))
}
