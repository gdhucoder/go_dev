package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/Shopify/sarama"
)

var (
	wg sync.WaitGroup
)

func initKafka() {

	// localhost
	consumer, err := sarama.NewConsumer(strings.Split("127.0.0.1:9092", ","), nil)
	if err != nil {
		fmt.Printf("Failed to start consumer: %s", err)
		return
	}
	partitionList, err := consumer.Partitions("test") // topic name: test
	if err != nil {
		fmt.Println("Failed to get the list of partitions: ", err)
		return
	}
	fmt.Println(partitionList)             // according to partition
	for partition := range partitionList { // each partition
		pc, err := consumer.ConsumePartition("test", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("Failed to start consumer for partition %d: %s\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		go func(pc sarama.PartitionConsumer) {
			wg.Add(1) // wait group goroutine
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				fmt.Println()
			}
			wg.Done()
		}(pc)
	}
	time.Sleep(time.Hour) // wait,then get message
	wg.Wait()
	consumer.Close()
}
