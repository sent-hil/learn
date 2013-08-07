package main

import "fmt"

func main() {
	hash := map[string]interface{}{
		"one": "one",
	}

	if result, ok := hash["one"]; ok {
		fmt.Println(result)
	}

	// throws error
	fmt.Println(result)

	result, ok := hash["one"]
	if ok {
		fmt.Println(result)
	}

	fmt.Println(result)
}
