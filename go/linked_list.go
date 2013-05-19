package main

import (
	"fmt"
)

type LinkedList struct {
	value int
	next  *LinkedList
}

func main() {
	list1 := LinkedList{value: 0}
	list2 := LinkedList{value: 1, next: &list1}
	list3 := LinkedList{value: 2, next: &list2}
	list4 := LinkedList{value: 3, next: &list3}

	iter_through_list(list4)
}
