package main

import (
	"fmt"
	"time"
)

func main() {
	if time.Now().Hour() < 12 {
		fmt.Println("Good Morning")
	} else {
		fmt.Println("Good afternoon")
	}

	birthday, _ := time.Parse("Jan 2 2006", "Nov 10 2009")
	age := time.Since(birthday)
	fmt.Printf("Go is %d days old\n", age/(time.Hour*24))

	t := time.Now()
	fmt.Println(t.In(time.UTC))

	home, _ := time.LoadLocation("Australia/Sydney")
	fmt.Println(t.In(home))
}
