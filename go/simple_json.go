package main

import (
	"fmt"
	simplejson "github.com/bitly/go-simplejson"
)

func main() {
	data := []byte(`{"greeting":{"hello":"world"}}`)

	js, err := simplejson.NewJson(data)
	if err != nil {
		fmt.Println("Got error", err)
	}
	fmt.Println(js.Get("greeting").Get("hello"))
	fmt.Println(js.GetPath("greeting", "hello"))
}
