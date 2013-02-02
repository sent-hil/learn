package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0

	for _, num := range nums {
		sum += num
	}

	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index of 3:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}

	channel := make(chan int, 2)
	channel <- 1
	channel <- 2

	//need to close channel, before you can iterate over it
	close(channel)

	for msg := range channel {
		fmt.Println(msg)
	}
}
