package main

import (
	"fmt"
	"strings"
)

var botName = "koding_bot"

func main() {
	var messageArgs = []string{"#koding-test", "koding_bot: gist.github.com/A"}

	if strings.HasPrefix(messageArgs[1], botName) {
		var withoutBotName = strings.SplitAfter(messageArgs[1], ": ")
		fmt.Println(withoutBotName[1])

		if strings.Contains(withoutBotName[1], "gist.github.com") {
			var url = buildGistTeamworkImportUrl(withoutBotName[1])
			fmt.Println(url)
		}
	}
}

func buildGistTeamworkImportUrl(gistUrl string) string {
	return "https://koding.com/Teamwork?import=" + gistUrl
}
