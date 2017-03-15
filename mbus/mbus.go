package mbus

import (
	"fmt"
)

type BusType string

const (
	Kafka BusType = "kafka"
	Nats          = "nats"
)

type NewBusFun func(servers []string) (MsgBus, error)

var MsgBusMap map[BusType]NewBusFun

func init() {
	MsgBusMap = make(map[BusType]NewBusFun)
}

func RegistryBus(busType BusType, newBusFun NewBusFun) {
	if _, ok := MsgBusMap[busType]; !ok {
		MsgBusMap[busType] = newBusFun
	}
}

type MsgBus interface {
	CreateTopic(topicName string) (Topic, error)
	Close()
}

type Topic interface {
	Publish(jsonObj interface{}) error
	Subscribe(handleFunc Handler) error
	Unsubscribe() error
}

type Handler interface{}

func NewBus(busType BusType, servers []string) (mbus MsgBus, err error) {
	if newBusFun, ok := MsgBusMap[busType]; ok {
		return newBusFun(servers)
	}

	return nil, fmt.Errorf("The message bus type %s is not supported", busType)
}
