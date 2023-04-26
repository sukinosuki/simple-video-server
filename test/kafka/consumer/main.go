package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"time"
)

func main() {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		//"bootstrap.servers": "192.168.10.182:9092",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}

	err = consumer.SubscribeTopics([]string{"myTopic", "^aRegex.*[Tt]opic"}, nil)
	if err != nil {
		panic(err)
	}

	run := true
	for run {
		msg, err := consumer.ReadMessage(time.Second)
		if err != nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else if !err.(kafka.Error).IsTimeout() {
			// The client will automatically try to recover from all errors.
			// Timeout is not considered an error because it is raised by
			// ReadMessage in absence of messages.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	err = consumer.Close()
	if err != nil {
		fmt.Println("consumer关闭错误 ", err)
	}
}
