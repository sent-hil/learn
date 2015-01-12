package main

import (
	"log"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("koding").C("jNewStatusUpdates")

	result := []interface{}{}
	err = c.Find(bson.M{"hostKite": nil}).All(&result)
	if err != nil {
		log.Println(err)
	}

	objectId := result[0].(bson.M)["_id"].(bson.ObjectId)
	c.Update(bson.M{"_id": objectId}, bson.M{"$set": bson.M{"hostKite": nil}})

	log.Println(result[0])
	log.Println(result[0])
}
