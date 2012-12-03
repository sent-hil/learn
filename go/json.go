package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))
}
