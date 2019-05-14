package main

import (
	"context"

	"github.com/astaxie/beego/logs"
	elastic "github.com/olivere/elastic"
)

type LogMessage struct {
	App     string
	Topic   string
	Message string
}

var (
	esClient *elastic.Client
)

func initES(addr string) (err error) {
	client, errRt := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(addr))
	if errRt != nil {
		logs.Error("connect es error", err)
		err = errRt
		return
	}
	esClient = client
	logs.Debug("conn es succ")
	return
}

func sendToES(topic string, data []byte) (err error) {
	ctx := context.Background()
	msg := &LogMessage{}
	msg.Topic = topic
	msg.Message = string(data)
	_, err = esClient.Index().Index(topic).Type(topic).BodyJson(msg).Do(ctx)
	if err != nil {
		panic(err)
		return
	}
	return
}
