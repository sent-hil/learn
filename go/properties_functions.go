package main

import (
	"log"
)

type me func()

func (meptr me) Hello() func([]map[string]interface{}) map[string]interface{} {
	type data map[string]interface{}

	var d = make(data)

	return func(values []map[string]interface{}) map[string]interface{} {
		for _, item := range values {
			for k, v := range item {
				log.Println(k, v)

				d[k] = v
			}
		}

		return d
	}
}

func main() {
	var m me
	var inside = make(map[string]interface{})
	inside["a"] = "b"

	var mm = make([]map[string]interface{}, 1)
	mm[0] = inside

	log.Println(mm)

	a := m.Hello()(mm)
	a["c"] = "d"

	log.Println(a)
}
