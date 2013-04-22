package main

import (
	"container/heap"
	"log"
	"math/rand"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(v interface{}) {
	a := *h
	a = append(a, v.(int))
	*h = a
}

func (h *IntHeap) Pop() interface{} {
	a := *h
	n := len(a)
	v := a[n-1]
	*h = a[0 : n-1]
	return v
}

func main() {
	h := make(IntHeap, 0)
	log.Printf("%v", h)
	for i := 0; i < 10; i++ {
		num := rand.Intn(25)
		heap.Push(&h, num)
	}
	log.Printf("%v", h)

	l := h.Len()
	ints := make([]int, 0, 1)
	for i := 0; i < l; i++ {
		ints = append(ints, heap.Pop(&h).(int))
	}

	log.Printf("%v", ints)
	log.Printf("%v", h)
}
