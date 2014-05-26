package main

import (
	"log"
	"time"
)

func main() {
	log.Println((time.Hour * 24).Hours())

}

func getTodayDate() time.Time {
	year, month, day := time.Now().Date()
	return time.Date(year, month, day, 0, 0, 0, 0, currentTimeLocation)
}
