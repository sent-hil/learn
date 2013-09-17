package main

import (
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	var timeout = time.Duration(2 * time.Second)

	dialTimeout := func(network, addr string) (net.Conn, error) {
		return net.DialTimeout(network, addr, timeout)
	}

	req, err := http.NewRequest("POST", "http://localhost:8080", nil)
	if err != nil {
		log.Println(err)
	}

	transport := http.Transport{
		Dial: dialTimeout,
	}

	client := http.Client{
		Transport: &transport,
	}

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	log.Println(res)
}
