package main

import (
	"fmt"
	"go_dev/day11/logagent/kafka"
	"go_dev/day11/logagent/tailf"
	"time"
)

func runServer() {
	fmt.Println("server is running.")

	for {
		msg := tailf.GetOneLine() // read from log
		fmt.Printf("read by tailf, sending to kafka -> msg: %s, topic: %s\n", msg.Msg, msg.Topic)
		kafka.SendToKafka(msg.Msg, msg.Topic)
		time.Sleep(time.Second * 5)
	}
	// return
}
