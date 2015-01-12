package main

import (
	"fmt"
	"reflect"
	"runtime"
)

type A struct{}

func (a *A) Hello() {}

func main() {
	a := A{}
	rfn := reflect.ValueOf(a.Hello)

	fmt.Println(runtime.FuncForPC(rfn.Pointer()).Name())
}
