package main

import (
	"fmt"
	"os/exec"
	"unicode/utf8"
)

func main() {
	result, _ := exec.Command("cat", "chan1.go").Output()
	//result := "\U000065e5\U0000672c\U00008a9e"

	for _, str := range result {
		st := []byte{str}
		a, b := utf8.DecodeRune(st)

		if b != 1 {
			fmt.Println(st, a, b)
		}
	}
}
