package main

import "fmt"

type Query map[string]string

func (q Query) get(s string) string {
	return q[s]
}

func main() {
	vals := make(map[string]string)
	q := Query{}

	for k, v := range vals {
		q[k] = v
	}

	q["a"] = "a"
	st := "a"

	fmt.Println(q.get(st))
}
