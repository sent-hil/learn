package main

import (
	"fmt"
	"time"
)

func main() {
	//creates channel that will current time after duration
	//specified
	timer1 := time.NewTimer(time.Millisecond * 100)

	<-timer1.C
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Millisecond * 200)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()

	//advantage of using a timer over timeout is so you can
	//stop is manually if necessary
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
