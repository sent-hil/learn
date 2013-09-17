package main

import (
	"log"

	"time"
)

func main() {
	timer := time.NewTimer(1 * time.Second)

	go func() {
		<-timer.C
		panic("req too forever")
	}()

	time.Sleep(5 * time.Second)

	defer func() {
		log.Println("exited")
		timer.Stop()
	}()
}
