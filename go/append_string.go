package main

import (
	"fmt"
	"time"
)

func main() {
	a := fmt.Sprintf("%v", time.Now())
	fmt.Printf("%T, %v", a, a)
}
