package main

import "fmt"

func main() {
	//make slice with initial size of 3
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	//append elems to it
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl3:", l)

	l = s[2:]
	fmt.Println("sl2:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	var array [5]int

	//create slice from array above
	slice := array[0:3]

	fmt.Println(len(slice) == 3)
	fmt.Println(cap(slice) == 5)

	//slices are just references, so any changes made to
	//original array will be visible here
	array[0] = 1
	fmt.Println(slice[0] == 1)

	res := append(slice, 5)
	fmt.Println(slice, res)
}
