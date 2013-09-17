package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	lastSeenMessage := time.Now()

	checker := time.Tick(1 * time.Second)
	producer := time.Tick(10 * time.Second)

	log.Println("Started")

	go func() {
		for _ = range producer {
			lastSeenMessage = time.Now()
			log.Println(lastSeenMessage)
		}
	}()

	for _ = range checker {
		checkTime := 5 * time.Second
		lastSeenMessageTime := lastSeenMessage.Add(checkTime).Unix()
		currentTime := time.Now().Unix()

		check := lastSeenMessageTime < currentTime

		log.Println(currentTime, lastSeenMessageTime, check)

		if check {
			panic(fmt.Sprintf("no message received in last %v", checkTime))
		}
	}
}
