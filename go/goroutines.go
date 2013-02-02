package main

import (
	"fmt"
	"time"
)

//mimics long running time by sleeping
func longCalculation(x int) int {
	time.Sleep(200)
	return 1
}

func main() {
	c := make(chan int)

	//wait for return in a goroutine
	go func(a int, c chan int) {
		result := longCalculation(a)
		c <- result
	}(17, c)

	//goroutine are nonblocking, so this will print before
	fmt.Println("non blocking")

	//waits for return from channel
	x := <-c
	fmt.Println(x)
}
