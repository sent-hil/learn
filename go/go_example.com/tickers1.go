package main

import (
	"fmt"
)

func main() {
	channel := make(chan int)

	//iterate over msgs in channel
	go func() {
		for t := range channel {
			fmt.Println("Ticker at", t)
		}
	}()

	channel <- 1
	channel <- 2
}
