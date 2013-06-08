package main

import (
	"log"
)

func main() {
	multipleArguments()
	multipleArguments(map[string]interface{}{"a": "b"})
}

func multipleArguments(options ...map[string]interface{}) {
	option := options[0]
	log.Println(option)
}
