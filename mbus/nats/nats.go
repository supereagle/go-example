package nats

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/apcera/nats"
	"github.com/supereagle/go-example/mbus"
)

const (
	DefaultMaxReconnect  = 5
	DefaultReconnectWait = 2
)

func init() {
	newBusFun := func(servers []string) (mbus.MsgBus, error) {
		if servers == nil || len(servers) <= 0 {
			return nil, fmt.Errorf("The servers must not be empty.")
		}
		opts := nats.DefaultOptions
		opts.Servers = servers
		opts.MaxReconnect = DefaultMaxReconnect

		opts.ReconnectWait = time.Duration(DefaultReconnectWait) * time.Second

		opts.NoRandomize = true
		nc, err := opts.Connect()
		if err != nil {
			return nil, fmt.Errorf("Fail to connect to the servers %s as %s", servers, err.Error())
		}
		enc, err := nats.NewEncodedConn(nc, "json")
		if err != nil {
			return nil, fmt.Errorf("Fail to new encoded connection as %s", err.Error())
		}
		nats := &Nats{
			opts,
			enc,
		}
		return nats, nil
	}

	mbus.RegistryBus(mbus.Nats, newBusFun)
}

type Nats struct {
	opts     nats.Options
	netConnt *nats.EncodedConn
}

func NewNats(servers []string, maxReconnect, reconnectWaitSecond int) (mbus *Nats, err error) {
	if servers == nil || len(servers) <= 0 {
		err = errors.New("The servers must not be empty.") // i18n
		return
	}
	opts := nats.DefaultOptions
	opts.Servers = servers
	if maxReconnect > 0 {
		opts.MaxReconnect = maxReconnect
	} else {
		opts.MaxReconnect = DefaultMaxReconnect
	}
	if reconnectWaitSecond > 0 {
		opts.ReconnectWait = time.Duration(reconnectWaitSecond) * time.Second
	} else {
		opts.ReconnectWait = time.Duration(DefaultReconnectWait) * time.Second
	}
	opts.NoRandomize = true
	nc, err := opts.Connect()
	if err != nil {
		return
	}
	enc, err := nats.NewEncodedConn(nc, "json")
	if err != nil {
		return
	}
	mbus = &Nats{
		opts,
		enc,
	}
	return
}

func (mbus *Nats) CreateTopic(topicName string) (topic mbus.Topic, err error) {
	if mbus.netConnt.Conn.IsClosed() {
		err = errors.New("The message bus connection is closed.") // i18n
		return
	}
	if strings.TrimSpace(topicName) == "" {
		err = errors.New("The topic name must not be empty.") //18n
		return
	}

	topic = &Topic{
		mbus,
		strings.TrimSpace(topicName),
		nil,
	}
	return
}

func (mbus *Nats) Close() {
	defer mbus.netConnt.Close()
	mbus.netConnt.Flush()
}

type Topic struct {
	mbus *Nats
	name string
	sub  *nats.Subscription
}

func (tpc *Topic) Publish(jsonObj interface{}) (err error) {
	err = tpc.mbus.netConnt.Publish(tpc.name, jsonObj)
	if err != nil {
		return
	}
	err = tpc.mbus.netConnt.Flush()
	return
}

func (tpc *Topic) Subscribe(handleFunc mbus.Handler) (err error) {
	sub, err := tpc.mbus.netConnt.Subscribe(tpc.name, nats.Handler(handleFunc))
	if err != nil {
		return
	}
	tpc.sub = sub
	return
}

func (tpc *Topic) Unsubscribe() (err error) {
	if tpc.sub != nil {
		err = tpc.sub.Unsubscribe()
	}
	return
}
