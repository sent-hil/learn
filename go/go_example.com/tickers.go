package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	//create new channel that will send time with period
	//specified by duration argument
	ticker := time.NewTicker(time.Millisecond * 100)

	//iterate over ticker
	go func() {
		for t := range ticker.C {
			fmt.Println("Ticker at", t)
		}
	}()

	//sleep so goroutine can work
	time.Sleep(time.Millisecond * 500)

	ticker.Stop()
	fmt.Println("Ticker stopped")
}
