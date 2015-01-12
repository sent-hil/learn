package main

import (
	"log"
	"runtime/debug"

	"time"
)

func main() {
	timer := time.NewTimer(1 * time.Second)

	go func() {
		<-timer.C
		panic("req too forever")
	}()

	debug.PrintStack()

	time.Sleep(5 * time.Second)

	defer func() {
		log.Println("exited")
		timer.Stop()
	}()
}
