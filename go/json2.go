package main

import (
	"encoding/json"
	"log"
	"time"
)

func main() {
	blob := `{"_id":"5199c7058f61295a3d000005","counts":{"followers":1}, "meta":{"createdAt":"2013-05-20T06:47:33.000Z"}}`
	var result map[string]*json.RawMessage
	json.Unmarshal([]byte(blob), &result)

	var id string
	json.Unmarshal(*result["_id"], &id)
	log.Printf("_id: %s", id)

	var counts map[string]interface{}
	json.Unmarshal(*result["counts"], &counts)
	log.Printf("counts: %v", counts["followers"])

	var meta map[string]interface{}
	json.Unmarshal(*result["meta"], &meta)

	var t time.Time
	json.Unmarshal(*meta, &t)
	log.Println(t)
}
