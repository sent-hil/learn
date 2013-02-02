package main

import "fmt"

type TZ int

const UTC TZ = 0 * 60 * 60
const EST TZ = -5 * 60 * 60

// define type which is a function
type Op func(int, int) int

type RPC struct {
	a, b   int
	op     Op
	result *int
}

func add(a, b int) int {
	return a + b
}

func main() {
	//define slice which holds strings
	weekend := []string{"Saturday", "Sunday"}

	//iterate over it, print index & name
	for index, day := range weekend {
		fmt.Println(index, day)
	}

	//append since it's a slice
	fmt.Println(append(weekend, "Monday"))

	//define hash where strings map to TZ, type defined above
	timeZones := map[string]TZ{
		"UTC": UTC, "EST": EST,
	}

	//iterate over it, print key & value
	for key, value := range timeZones {
		fmt.Println(key, value)
	}

	//create new instance of RPC type
	//think of it like Rpc.new(1,2..)
	//
	//notice add is a function that's passed as an value
	rpc := RPC{1, 2, add, new(int)}
	fmt.Println(rpc.op(rpc.a, rpc.b))
}
