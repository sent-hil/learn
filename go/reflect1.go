package main

import (
	"log"
	"reflect"
	"runtime"
)

func Hello() int {
	log.Println("hello")
	return 1
}

func main() {
	var hola = func() func() int {
		return Hello
	}

	callMe(hola)
}

func callMe(fn interface{}) {
	rfn := reflect.ValueOf(fn)

	log.Println(rfn.Pointer())
	log.Println(runtime.FuncForPC(rfn.Pointer()).Name())

	// result := make([]byte, 300)
	// runtime.Stack(result, true)

	// log.Println(string(result))
}
