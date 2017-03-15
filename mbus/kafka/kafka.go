package kafka

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/optiopay/kafka"
	"github.com/optiopay/kafka/proto"
	"github.com/supereagle/go-example/mbus"
)

const (
	// ClientId The default Kafka client id for PaaS platform
	ClientId = "paas-client"
)

func init() {
	newBusFun := func(servers []string) (mbus.MsgBus, error) {
		if servers == nil || len(servers) <= 0 {
			return nil, fmt.Errorf("The servers must not be empty.")
		}

		// Broker config
		conf := kafka.NewBrokerConf(ClientId)
		conf.AllowTopicCreation = false

		// Connect to kafka cluster
		broker, err := kafka.Dial(servers, conf)
		if err != nil {
			return nil, fmt.Errorf("Fail to connect to kafka cluster: %s", err)
		}

		kafka := &KafkaClient{broker}

		return kafka, nil
	}

	mbus.RegistryBus(mbus.Kafka, newBusFun)
}

var emptyMsgType = reflect.TypeOf(&proto.Message{})

// KafkaClient The Kafka client connected to any one of kafka servers
type KafkaClient struct {
	broker *kafka.Broker
}

// NewKafkaClient Creates the Kafka client, remember to close it after used.
func NewKafkaClient(servers []string) (*KafkaClient, error) {
	if servers == nil || len(servers) <= 0 {
		return nil, fmt.Errorf("The servers must not be empty.")
	}

	// Broker config
	conf := kafka.NewBrokerConf(ClientId)
	conf.AllowTopicCreation = false

	// Connect to kafka cluster
	broker, err := kafka.Dial(servers, conf)
	if err != nil {
		return nil, fmt.Errorf("Fail to connect to kafka cluster: %s", err)
	}

	kafka := &KafkaClient{broker}

	return kafka, nil
}

// CreateTopic Creates the topic with the specified name
func (k *KafkaClient) CreateTopic(name string) (mbus.Topic, error) {
	// Get the partitions of the topic
	md, err := k.broker.Metadata()
	if err != nil {
		return nil, fmt.Errorf("Fail to get the metadata of the Kafka client")
	}

	partitionIds := []int32{}
	for _, mt := range md.Topics {
		if mt.Name == name {
			for _, partition := range mt.Partitions {
				partitionIds = append(partitionIds, partition.ID)
			}
			break
		}
	}

	numPartitions := int32(len(partitionIds))
	if numPartitions == 0 {
		// TODO (robin) Create the topic with default config when not exists
		return nil, fmt.Errorf("The topic %s does not exist", name)

		// Three partition for new created topics in default
		// partitionIds = []int32{0, 1, 2}
		// numPartitions = 3
	}

	consumers := make([]kafka.Consumer, numPartitions)
	for i, pId := range partitionIds {
		conf := kafka.NewConsumerConf(name, pId)
		conf.StartOffset = kafka.StartOffsetNewest
		consumer, err := k.broker.Consumer(conf)
		if err != nil {
			return nil, fmt.Errorf("Fail to create kafka consumer for topic %s with partition %d as %s", name, pId, err.Error())
		}
		consumers[i] = consumer
	}

	mx := kafka.Merge(consumers...)

	pConf := kafka.NewProducerConf()
	producer := k.broker.Producer(pConf)

	distributer := kafka.NewRoundRobinProducer(producer, numPartitions)

	topic := &Topic{
		name:        name,
		mx:          mx,
		distributer: distributer,
	}

	return topic, nil
}

// Close Closes the Kafka client by closing its broker
func (k *KafkaClient) Close() {
	k.broker.Close()
}

// Topic The topic can produce and consume messages
type Topic struct {
	name        string
	mx          *kafka.Mx
	distributer kafka.DistributingProducer
}

// Publish Publishes the messages for the topics
func (t *Topic) Publish(jsonObj interface{}) error {
	bs, err := json.Marshal(jsonObj)
	if err != nil {
		return fmt.Errorf("Fail to marshal the json object to be produced for the topic %s", t.name)
	}

	msg := &proto.Message{Value: bs}

	if _, err := t.distributer.Distribute(t.name, msg); err != nil {
		return fmt.Errorf("Fail to produce messages for the topic %s", t.name)
	}

	return nil
}

// Subscribe Registries a handler to consume the messages from the topics
func (t *Topic) Subscribe(handleFunc mbus.Handler) (err error) {
	argType, numArgs := argInfo(handleFunc)
	cbValue := reflect.ValueOf(handleFunc)
	wantsRaw := (argType == emptyMsgType)

	if numArgs != 1 {
		return fmt.Errorf("Now just support one parameters in handler")
	}

	go func() {
		for {
			msg, err := t.mx.Consume()
			if err != nil {
				// TODO Add the error handler
				fmt.Println(err)
				if err == kafka.ErrMxClosed {
					return
				}
				continue
			}

			var oV []reflect.Value
			if wantsRaw {
				oV = []reflect.Value{reflect.ValueOf(msg)}
			} else {
				var oPtr reflect.Value
				if argType.Kind() != reflect.Ptr {
					oPtr = reflect.New(argType)
				} else {
					oPtr = reflect.New(argType.Elem())
				}

				err := decode(msg.Value, oPtr.Interface())
				if err != nil {
					// TODO Add the error handler
					fmt.Println(err)
					continue
				}
				if argType.Kind() != reflect.Ptr {
					oPtr = reflect.Indirect(oPtr)
				}

				oV = []reflect.Value{oPtr}
			}

			cbValue.Call(oV)
		}
	}()

	return nil
}

func (t *Topic) Unsubscribe() error {
	t.mx.Close()

	return nil
}

// Dissect the cb Handler's signature
func argInfo(cb interface{}) (reflect.Type, int) {
	cbType := reflect.TypeOf(cb)
	if cbType.Kind() != reflect.Func {
		panic("nats: Handler needs to be a func")
	}
	numArgs := cbType.NumIn()
	if numArgs == 0 {
		return nil, numArgs
	}
	return cbType.In(numArgs - 1), numArgs
}

func decode(data []byte, vPtr interface{}) (err error) {
	switch arg := vPtr.(type) {
	case *string:
		// If they want a string and it is a JSON string, strip quotes
		// This allows someone to send a struct but receive as a plain string
		// This cast should be efficient for Go 1.3 and beyond.
		str := string(data)
		if strings.HasPrefix(str, `"`) && strings.HasSuffix(str, `"`) {
			*arg = str[1 : len(str)-1]
		} else {
			*arg = str
		}
	case *[]byte:
		*arg = data
	default:
		err = json.Unmarshal(data, arg)
	}
	return
}
