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

var totalMsgs = 10000

func produceProtos(p *kafka.Producer, count int) (time.Time, time.Time) {
	// Produce messages to topic (asynchronously)
	topic := "protobuf"

	// Build and initialize user message.
	msg := new(pbproto.MarkReq)
	msg.SlNo = 123
	msg.Name = "ProtoBuf"
	msg.Subject = pbproto.Subject_PHYSICS
	btes, _ := proto.Marshal(msg)

	startTime := time.Now()

	for i := 1; i <= count; i++ {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          btes,
		}, nil)
	}

	endTime := time.Now()

	return startTime, endTime
}

func produceStructs(p *kafka.Producer, count int) (time.Time, time.Time) {
	// Produce messages to topic (asynchronously)
	topic := "struct"

	// Build and initialize user message.
	msg := new(message.MarkReq)
	msg.SlNo = 123
	msg.Name = "StructMsg"
	msg.Subject = message.SubjectPHYSICS
	btes, _ := json.Marshal(msg)

	startTime := time.Now()

	for i := 1; i <= count; i++ {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          btes,
		}, nil)
	}

	endTime := time.Now()

	return startTime, endTime
}

func main() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	protoStartTime, protoEndTime := produceProtos(p, totalMsgs)
	structStartTime, structEndTime := produceStructs(p, totalMsgs)

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)

	fmt.Printf("Dealing with protobufs ==> Number of messages sent :%d  , Start Time : %v , End Time : %v\n\n", totalMsgs, protoStartTime, protoEndTime)
	fmt.Printf("Dealing with structs ==> Number of messages sent :%d  , Start Time : %v , End Time : %v\n\n", totalMsgs, structStartTime, structEndTime)

	// Wait couple of seconds for delivery reports
	time.Sleep(2 * time.Second)
}
