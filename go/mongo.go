package main

import (
	"io"
	"labix.org/v2/mgo"
	"log"
	"time"
)

func main() {
	session, err := mgo.Dial("localhost:27017?connect=direct")
	if err != nil {
		log.Println(err)
	}

	session.SetSafe(&mgo.Safe{})

	log.Println(session.Mode())

	defer session.Close()

	rels := session.DB("koding").C("relationships")

	c := time.Tick(2 * time.Second)
	for _ = range c {
		result := map[string]interface{}{}
		err = rels.Find(nil).One(&result)

		if err != nil {
			if err == io.EOF {
				session.Refresh()
			} else {
				return
			}
		}

		log.Println(result)
	}

	log.Println(session)
}
