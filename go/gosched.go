package main

import (
	"fmt"
	"runtime"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

func main() {
	runtime.GOMAXPROCS(2)
	go say("world")
	say("hello")

	//without this line "world" will never be printed since
	//the process ends before
	time.Sleep(time.Second * 1)
}
