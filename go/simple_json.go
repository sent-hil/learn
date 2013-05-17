package main

import (
	simplejson "github.com/bitly/go-simplejson"
	"log"
)

func main() {
	blob := `{"Title":"Crusade", "Actor":["Ford", "Sean"]}`
	json, err := simplejson.NewJson([]byte(blob))
	if err != nil {
		log.Fatal(err)
	}
	title, _ := json.Get("Title").String()
	log.Println(title)

	actorsList, exists := json.CheckGet("Actor")
	log.Println(actorsList, exists)

	actors, err := actorsList.Array()
	if err != nil {
		log.Fatal(err)
	}

	for _, actor := range actors {
		log.Printf("%s acted in %s", actor, title)
	}
}
