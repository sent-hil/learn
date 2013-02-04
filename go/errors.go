package main

import (
	"errors"
	"fmt"
)

func greater_than_ten(value int) (bool, error) {
	if value <= 10 {
		return false, errors.New("not true")
	}

	return true, nil
}

func main() {
	var result bool
	var err error

	result, err = greater_than_ten(10)
	fmt.Println("10 greater than 10: ", result)
	fmt.Println("10 greater than 10, error: ", err)

	result, err = greater_than_ten(11)
	fmt.Println("11 less than 10: ", result)
	fmt.Println("11 less than 10, error: ", err)
}
