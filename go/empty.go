package main

import "log"

func main() {
	a := make(map[string]interface{})

	log.Println(a["a"] == nil)

	b, ok := a["a"]
	if !ok {
		b = make(map[string]interface{})
		a["a"] = b
	}

	log.Println(b)

  if b
}
