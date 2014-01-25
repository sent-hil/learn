package main

import (
	"log"
	"net/http"
	"net/url"
)

const graphityEventsUrl = "http://localhost:7474/graphity/events"

func main() {
	req, err := http.NewRequest("POST", graphityEventsUrl, nil)
	if err != nil {
		panic(err)
	}

	body := map[string]string{
		"source": "http://127.0.0.1:7474/db/data/node/27255",
		"event":  "http://127.0.0.1:7474/db/data/node/27256",
	}

	req.PostForm = url.Values{}
	for k, v := range body {
		req.PostForm.Add(k, v)
	}
	req.PostForm.Encode()

	req.Form = url.Values{}
	for k, v := range body {
		req.Form.Add(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	log.Println(resp)
}
