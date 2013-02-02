package main

import "fmt"

func main() {
	//make Hash of {string => int}
	m := make(map[string]int)

	//Hash with defaults
	mo := map[string]int{"foo": 1, "bar": 2}
	var mm = map[string]int{"foo": 1, "bar": 2}

	var n map[string]int

	fmt.Println(n)

	//setters
	m["k1"] = 7
	m["k2"] = 13

	//getters
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("map:", m)
	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	//returns false if key doesn't exist
	_, prs := m["k2"]
	fmt.Println("prs:", prs)
	fmt.Println("map:", mm, mo)
}
