package main

import (
	"flag"
	"fmt"
	"time"
)

//flag to be parsed, default message, help message
var message = flag.String("message", "Hello!", "what to say")
var delay = flag.Duration("delay", 1*time.Second, "how long to wait")

func main() {
	flag.Parse()

	//message is a pointer
	fmt.Println(*message, *delay)
	time.Sleep(*delay)
}
