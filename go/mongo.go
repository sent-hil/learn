package main

import (
	"labix.org/v2/mgo"
	"log"
)

func main() {
	conn, _ = mgo.Dial(MONGO_CONN_STRING)
	log.Println(mgo)
}
