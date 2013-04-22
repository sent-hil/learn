package main

import (
	"fmt"
	"time"
)

// Similar to setInterval, c fires every 1 second
func main() {
	c := time.Tick(1 * time.Second)
	fmt.Println("before")
	for tm := range c {
		fmt.Println(tm) //print current time
	}

	// blocking, never fired
	fmt.Println("after")
}
