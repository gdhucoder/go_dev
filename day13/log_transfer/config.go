package main

import (
	"fmt"

	"github.com/astaxie/beego/config"
)

type LogConfig struct {
	KafkaAddr string
	ESAddr    string
	LogPath   string
	LogLevel  string
	Topic     string
}

var (
	logConfig *LogConfig
)

func initConfig(configType, configPath string) (err error) {
	conf, err := config.NewConfig(configType, configPath)
	if err != nil {
		fmt.Println("new config failed, err:", err)
		return
	}

	logConfig = &LogConfig{}

	logPath := conf.String("log::log_path")
	logConfig.LogPath = logPath

	logLevel := conf.String("log::log_level")
	logConfig.LogLevel = logLevel

	kafkaAddr := conf.String("kafka::kafka_addr")
	logConfig.KafkaAddr = kafkaAddr

	topic := conf.String("kafka::topic")
	logConfig.Topic = topic

	esAddr := conf.String("es::es_addr")
	logConfig.ESAddr = esAddr

	return
}
