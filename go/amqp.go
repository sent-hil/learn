package main

import (
	"fmt"
	"streadway/amqp"
)

func main() {
	connection, err := amqp.Dial(os.Getenv("AMQP_URL"))
}
