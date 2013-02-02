package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	//initialize
	fmt.Println(Person{"Bob", 20})
	fmt.Println(Person{"Alice", 30})

	//initialize with keyword vars
	fmt.Println(Person{name: "Fred"})
	fmt.Println(&Person{name: "Ann", age: 40})

	s := Person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
