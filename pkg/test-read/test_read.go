package main

import (
	"Kaf-Protobuf/pkg/message"
	"Kaf-Protobuf/pkg/pbproto"
	"confluent-kafka-go/kafka"
	"fmt"
	"time"

	"encoding/json"

	"github.com/golang/protobuf/proto"
)

func readMsgs(c *kafka.Consumer) {

	var pbMsg pbproto.MarkReq
	var structMsg = new(message.MarkReq)

	var topicProtobuf = "protobuf"

	var protoTime = time.Now()
	var structTime = time.Now()

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			if *msg.TopicPartition.Topic == topicProtobuf {
				proto.Unmarshal(msg.Value, &pbMsg)
				fmt.Printf("Partition  : %v , Message Received at 'protobuf' topic : %v at time : %v\n", *msg.TopicPartition.Topic, pbMsg, time.Now())
				protoTime = time.Now()
			} else {
				json.Unmarshal(msg.Value, structMsg)
				fmt.Printf("Partition  : %v , Message Received at 'struct' topic : %v at time : %v\n", *msg.TopicPartition.Topic, structMsg, time.Now())
				structTime = time.Now()
			}

			fmt.Printf("\nLast message processed time == Protos : %v  ,  Structs : %v \n", protoTime, structTime)
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}

func main() {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"struct", "protobuf", "^aRegex.*[Tt]opic"}, nil)

	readMsgs(c)

	c.Close()
}
