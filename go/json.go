package main

import (
	"encoding/json"
	"fmt"
	"unicode/utf8"
)

func main() {
	strings := []interface{}{"apple", "peach", "pear", 0xffff}
	jsons, _ := json.Marshal(strings)
	fmt.Println(jsons)

	fmt.Println(utf8.ValidString(0xU00110000))
}
