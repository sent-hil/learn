package main

import (
	"fmt"
	"os"
	"runtime/pprof"
)

type LinkedList struct {
	value int
	next  *LinkedList
}

func main() {
	f, _ := os.Create("ll.prof")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	list1 := LinkedList{value: 0}
	list2 := LinkedList{value: 1, next: &list1}
	list3 := LinkedList{value: 2, next: &list2}
	list4 := LinkedList{value: 3, next: &list3}

	iter_through_list(list4)
}

func iter_through_list(list LinkedList) {
	fmt.Println("Value: ", list.value)

	if list.next != nil {
		iter_through_list(*list.next)
	} else {
		fmt.Println("At end of list.")
	}
}
