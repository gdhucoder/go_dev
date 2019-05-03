package conf

import (
	"fmt"

	"github.com/astaxie/beego/config"
)

type Config struct {
	LogLevel  string
	LogPath   string
	KafkaAddr string
	Collector []Collector
}

type Collector struct {
	LogPath  string
	Topic    string
	ChanSize int
}

var (
	AppConfig *Config
)

func InitConf(configType, configPath string) (err error) {
	conf, err := config.NewConfig(configType, configPath)
	if err != nil {
		fmt.Println("new config failed, err:", err)
		return
	}

	logPath := conf.String("server::log_path")
	if len(logPath) == 0 {
		panic("log path is not defined!")
	}
	AppConfig = &Config{}
	AppConfig.LogPath = logPath
	logLevel := conf.String("server::log_level")
	AppConfig.LogLevel = logLevel
	kafkaAddr := conf.String("kafka::kafka_addr")
	AppConfig.KafkaAddr = kafkaAddr
	loadCollector(conf)
	return
}

func loadCollector(conf config.Configer) (err error) {
	collcetor := Collector{}
	chanSize, err := conf.Int("collect::chan_size")
	if err != nil {
		chanSize = 100
	}
	collcetor.ChanSize = chanSize
	collcetor.LogPath = conf.String("collect::log_path")
	collcetor.Topic = conf.String("collect::topic")
	AppConfig.Collector = append(AppConfig.Collector, collcetor)
	return
}
