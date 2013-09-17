package main

import (
	"fmt"
	"koding/databases/neo4j"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	for i := 0; i < 1000; i++ {
		log.Println(i)

		var sourceId = "1"
		var sourceName = "guest"
		var sourceContent = "{}"

		sourceNode := neo4j.CreateUniqueNode(sourceId, sourceName)
		neo4j.UpdateNode(sourceId, sourceContent)

		rand.Seed(time.Now().UTC().UnixNano())
		var targetId = strconv.Itoa(rand.Int())
		var targetContent = "{}"

		targetNode := neo4j.CreateUniqueNode(targetId, targetId)
		neo4j.UpdateNode(targetId, targetContent)

		source := fmt.Sprintf("%s", sourceNode["create_relationship"])
		target := fmt.Sprintf("%s", targetNode["self"])

		neo4j.CreateRelationshipWithData("member", source, target, "{}")
	}
}
