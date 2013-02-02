package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 5
	v := reflect.ValueOf(x)
	fmt.Println("val:", x)
	fmt.Println("float?:", v.Kind() == reflect.Float64)
}
