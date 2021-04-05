package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

var kafkaServerPool []string = []string{
	"benjamin142857.ticp.vip:55077",
}

func testGetTopicList() {
	config := sarama.NewConfig()

	client, err := sarama.NewClient(kafkaServerPool, config)
	if err != nil {
		fmt.Printf("metadata_test try create client err :%s\n", err.Error())
		return
	}
	defer func(){ _ = client.Close() }()


	topics, err := client.Topics()
	if err != nil {
		fmt.Printf("try get topics err %s\n", err.Error())
		return
	}

	for k, topic := range topics {
		fmt.Println(k, topic)
	}
}

func testProducer1() {
	// 相关配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	// 初始化 client 连接
	producer, err := sarama.NewSyncProducer(kafkaServerPool, config)
	if err != nil {
		fmt.Println("sarama.NewSyncProducer error: ", err)
		return
	}
	defer func(){ _ = producer.Close() }()

	// 构建消息体
	msg := &sarama.ProducerMessage{ Topic: "bjm-test", Value: sarama.StringEncoder("Hello Benjamin666") }

	// 发送
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println("producer.SendMessage error: ", err)
		return
	}
	fmt.Printf("partition: %v, offser: %v\n", partition, offset)
}

func main() {
	//testGetTopicList()
	testProducer1()
}
