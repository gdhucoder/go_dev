package main

import (
	"strings"
	"sync"

	"github.com/astaxie/beego/logs"

	"github.com/Shopify/sarama"
)

type KafkaClient struct {
	client sarama.Consumer
	addr   string
	topic  string
	wg     sync.WaitGroup
}

var (
	kafkaClient *KafkaClient
)

func initKafka(addr, topic string) (err error) {

	kafkaClient = &KafkaClient{}

	// localhost
	consumer, err := sarama.NewConsumer(strings.Split(addr, ","), nil)
	if err != nil {
		logs.Error("Failed to start kafka: %s", err)
		return
	}

	kafkaClient.client = consumer
	kafkaClient.addr = addr
	kafkaClient.topic = topic
	return
}
