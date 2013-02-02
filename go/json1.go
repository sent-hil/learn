package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

const blob = `[
  {"Title":"Indiana", "URL":"http", "Something":"something"},
  {"Title":"Jones"}
]`

type Item struct {
	Title     string
	URL       string
	Something string
}

func main() {
	var items []*Item
	json.NewDecoder(strings.NewReader(blob)).Decode(&items)

	for _, item := range items {
		fmt.Println(item.Title, item.URL, item.Something)
	}
}
