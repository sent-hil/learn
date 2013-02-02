package main

import "fmt"

func main() {
	//while
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	//regular for
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	//loop
	for {
		fmt.Println("loop")
		break
	}
}
