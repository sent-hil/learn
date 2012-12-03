package main

import (
	"fmt"
	"reflect"
)

func main () {
	var x float64 = 5
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Kind() == reflect.Float64)
}
