package main

import "fmt"

func main() {
	//make channel
	messages := make(chan string)

	//send msgs in a seperate goroutine
	go func() {
		messages <- "ping"
	}()

	//receive msg, this blocks till msg arrives
	msg := <-messages
	fmt.Println(msg)
}
