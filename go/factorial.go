package main

import "fmt"

func iter(n int) (result int) {
	if n == 0 {
		result = 1
	} else {
		result = n * iter(n-1)
	}

	return
}

func iter_fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * iter_fact(n-1)
}

func iter_hashed(n int, hash map[int]int) (result int) {
	value, exists := hash[n]

	if exists == false {
		result = n * iter_hashed(n-1, hash)
	} else {
		result = value
	}

	return
}

func main() {
	fmt.Println("Recursive Factorial:", iter(7))
	fmt.Println("Iterative Factorial:", iter_fact(7))

	var hash = make(map[int]int)
	hash[0] = 1

	fmt.Println("Hashed Factorial:", iter_hashed(7, hash))
}
