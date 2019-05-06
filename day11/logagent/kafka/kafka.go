package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)

var (
	client sarama.SyncProducer
)

func InitKafka(kafakAddr string) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	client, err = sarama.NewSyncProducer([]string{kafakAddr}, config)
	if err != nil {
		logs.Error("init producer close, err:", err)
		return
	}
	logs.Debug("init Kafka succ!")
	return
}

func SendToKafka(data, topic string) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		logs.Error("send message failed,", err)
		return
	}
	logs.Debug("send succ: pid:%v offset:%v topic:%s", pid, offset, topic)
	return
}
