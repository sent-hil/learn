package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	// parse 1_First_post.md into First post
	raw := "1_First_post.md"

	var nameWithoutExtension string
	splitFileNames := strings.Split(raw, ".")
	if len(splitFileNames) > 0 {
		nameWithoutExtension = splitFileNames[0]
	} else {
		nameWithoutExtension = ""
	}

	var joinedTitle string
	splitFileNames = strings.Split(nameWithoutExtension, "_")
	if len(splitFileNames) > 1 {
		for _, str := range splitFileNames[1:] {
			joinedTitle += fmt.Sprintf("%s ", str)
		}
	} else {
		joinedTitle = ""
	}

	log.Println(joinedTitle)
}
