package main

import "log"

func main() {
	data := map[string]interface{}{
		"bongo_": map[string]interface{}{
			"constructorName": "Relationship",
			"instanceId":      "1f2e30fd7ed58b138d7e35dbd75e712a",
		},
		"data": map[string]interface{}{
			"targetId":   "5200004c2531290000000003",
			"targetName": "JAccount",
			"sourceId":   "5200004c2531290000000002",
			"sourceName": "JUser",
			"as":         "owner",
			"_id":        "5200004c2531290000000004",
		},
	}

	bongo, ok := data["bongo_"].(map[string]interface{})
	if !ok {
		log.Println("no bongo found")
	}

	log.Println(bongo)

	rdata, ok := data["data"].(map[string]interface{})
	if !ok {
		log.Println("no data found")
	}

	log.Println(rdata)

	// will panic if no 'data'
	id, ok := data["data"].(map[string]interface{})["_id"].(string)
	if !ok {
		log.Println("no data found")
	}

	log.Println(id)
}
