package main

import (
	"fmt"
	"time"

	"github.com/supereagle/go-example/mbus"
	// Import the package for Kafka client when use mbus based on Kafka.
	//_ "github.com/supereagle/go-example/mbus/kafka"
	_ "github.com/supereagle/go-example/mbus/nats"
)

type person struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

var topic string = "foo"

var persons []person = []person{
	person{
		Name: "Robin",
		Age:  30,
	},
	person{
		Name: "Tom",
		Age:  20,
	},
}

// Demo for the usage of mbus.
func main() {
	// Use the offical test nats server for mbus based on nats.
	msgBus, err := mbus.NewBus(mbus.Nats, []string{"nats://demo.nats.io:4222"})
	// If want to use mbus based on Kafka, choose the Kafka mbus type and specify the Kafka servers.
	// msgBus, err := mbus.NewBus(mbus.Kafka, []string{"localhost:9092", "localhost:9093"})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer msgBus.Close()

	topic, err := msgBus.CreateTopic(topic)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if err := topic.Subscribe(handler); err != nil {
		fmt.Println(err.Error())
		return
	}
	defer func() {
		if err := topic.Unsubscribe(); err != nil {
			fmt.Println(err.Error())
			return
		}
	}()

	for _, person := range persons {
		if err := topic.Publish(person); err != nil {
			fmt.Println(err.Error())
			return
		}
		time.Sleep(1 * time.Second)
	}
}

// handler the handler for message.
func handler(p *person) {
	fmt.Printf("The person received: %+v\n", p)
}
