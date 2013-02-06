package main

import (
	"flag"
	"fmt"
	"github.com/streadway/amqp"
)

func main() {
	body := flag.String("m", "Hello from Koding!", "")
	flag.Parse()

	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	fmt.Println("Connection error: ", err)

	channel, err := connection.Channel()
	fmt.Println("Channel error: ", err)

	defer connection.Close()

	//err = channel.ExchangeDeclare("amq.direct", "direct", true, false, false, false, nil)
	//fmt.Println("Exchange error: ", err)

	msg := amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		Body:         []byte(*body),
	}

	err = channel.Publish("amq.direct", "routing_key", false,
		false, msg)
	fmt.Println("Publish error: ", err)
}
