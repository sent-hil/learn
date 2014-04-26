package main

import "fmt"
import "time"

func main() {
	sevenDaysAgo := time.Now().Sub(time.Hour * 24 * 7)
	fmt.Println(sevenDaysAgo)
}
