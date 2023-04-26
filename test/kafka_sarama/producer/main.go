package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()

	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	fmt.Println("开始连接")
	client, err := sarama.NewSyncProducer([]string{"192.168.10.182:9092"}, config)
	if err != nil {
		fmt.Println("producer closed, err: ", err)
		panic(err)
	}
	defer client.Close()

	fmt.Println("开始发送消息1")
	// 例一: 发送单个消息
	msg := &sarama.ProducerMessage{
		Topic: "test",
	}
	msg.Value = sarama.StringEncoder("hello")
	//msg.Topic = "web_log"
	//content := "hello 233"
	send(client, msg)

	// 发送多个消息
	for _, word := range []string{"Welcome11", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
		msg.Value = sarama.StringEncoder(word)
		send(client, msg)
	}
}

func send(client sarama.SyncProducer, msg *sarama.ProducerMessage) {
	//msg.Value = sarama.StringEncoder(content)

	//client.SendMessage(msg)返回值有3个，第一个是数据所在分区，第二个是数据的偏移量，第三个是错误信息，无错误返回nil
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err ", err)
		return
	}

	fmt.Printf("pid: %v, offset: %v \n", pid, offset)
}
