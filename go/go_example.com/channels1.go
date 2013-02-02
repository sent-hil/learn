package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("done")

	//send msg when done
	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)

	//block till worker is done
	<-done
}
