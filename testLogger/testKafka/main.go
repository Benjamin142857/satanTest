package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func testGetTopicList() {
	config := sarama.NewConfig()

	client, err := sarama.NewClient([]string{"192.168.3.121:9092"}, config)
	if err != nil {
		fmt.Printf("metadata_test try create client err :%s\n", err.Error())
		return
	}
	defer client.Close()


	topics, err := client.Topics()
	if err != nil {
		fmt.Printf("try get topics err %s\n", err.Error())
		return
	}

	for k, topic := range topics {
		fmt.Println(k, topic)
	}
}

func main() {
	testGetTopicList()
}
