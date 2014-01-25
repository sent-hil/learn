package main

import (
	"fmt"
	"koding/workers/neo4jfeeder/mongohelper"
	"labix.org/v2/mgo/bson"
	"log"
)

func main() {
	sourceId := "51990577e77f6c2321000002"
	sourceName := "JStatusUpdate"

	sourceContent, err := mongohelper.FetchContent(bson.ObjectIdHex(sourceId), sourceName)
	if err != nil {
		fmt.Println("sourceContent", err)
		return
	}

	log.Println(sourceContent)
}
