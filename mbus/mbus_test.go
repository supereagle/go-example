package mbus_test

import (
	"testing"

	"github.com/supereagle/go-example/mbus"
	_ "github.com/supereagle/go-example/mbus/kafka"
	_ "github.com/supereagle/go-example/mbus/nats"
)

func TestMsgBusMap(t *testing.T) {
	if len(mbus.MsgBusMap) != 2 {
		t.Errorf("Not all message bus are registried")
	}

	if _, ok := mbus.MsgBusMap[mbus.Kafka]; !ok {
		t.Errorf("Kafka message bus is not registried")
	}

	if _, ok := mbus.MsgBusMap[mbus.Kafka]; !ok {
		t.Errorf("Nats message bus is not registried")
	}
}
