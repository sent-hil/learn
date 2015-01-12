package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron"
)

func main() {
	c := cron.New()

	c.AddFunc("0 0-59/5 * * * *", func() {
		fmt.Println("every 5 minutes", time.Now())
	})

	c.AddFunc("0 1-59/5 * * * *", func() {
		fmt.Println("every 5 minutes", time.Now())
	})

	c.Start()

	select {}
}
