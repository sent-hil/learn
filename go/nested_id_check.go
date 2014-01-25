package main

import (
	"errors"
	"fmt"
)

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

	fmt.Println(nestedIdCheck(data, "bongo_", "instanceId"))
	fmt.Println(nestedIdCheck(data, "data", "_id"))
	fmt.Println(nestedIdCheck(data, "bongo2_", "instanceId"))
}

func nestedIdCheck(raw map[string]interface{}, fieldNames ...string) (string, error) {
	if len(fieldNames) == 1 {
		value, ok := raw[fieldNames[0]].(string)

		if !ok {
			return "", errors.New("field not found")
		}

		return value, nil
	}

	data, ok := raw[fieldNames[0]].(map[string]interface{})
	if !ok {
		return "", errors.New("field not found")
	}

	return nestedIdCheck(data, fieldNames[1:]...)
}
