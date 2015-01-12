package main

import (
	"log"

	"github.com/garyburd/redigo/redis"
	"github.com/kr/pretty"
)

func main() {
	client, err := redis.Dial("tcp", "192.168.59.103:6379")
	if err != nil {
		log.Fatal(err)
	}

	pretty.Println(client.Do("SET", "hola", "booom"))
	pretty.Println(redis.String(client.Do("GET", "hola")))
}
