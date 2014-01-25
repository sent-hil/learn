package main

import (
	"labix.org/v2/mgo"
	"log"
)

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("koding").C("relationships")

	count, err := c.Find("1").Count()
	if err != nil {
		panic(err)
	}

	log.Println(count)
}
