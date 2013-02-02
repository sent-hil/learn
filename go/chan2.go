package main

import (
	"fmt"
	"reflect"
)

func main() {
	//make channel of strings
	var c chan string
	c = make(chan string)

	//send msg to channel
	go func() {
		c <- "Hello"
	}()

	//print msg received from above
	fmt.Println(<-c)

	//make a channel that accepts channel of strings
	cc := make(chan chan string)
	go func() {
		//send original channel as msg to new channel
		cc <- c
	}()

	//receive original channel
	loc := <-cc
	fmt.Println(reflect.TypeOf(loc))

	//send msg to original channel
	go func() {
		loc <- "hi there"
	}()

	//print msg received from above
	fmt.Println(<-loc)

	fmt.Println("Original channel: ", reflect.TypeOf(c))
	fmt.Println("Channel of channels: ", reflect.TypeOf(cc))
	fmt.Println("Original channel sent as msg", reflect.TypeOf(loc))
}
