package conf

import (
	"fmt"

	"github.com/astaxie/beego/config"
)

type Config struct {
	LogLevel  string
	LogPath   string
	KafkaAddr string
	EtcdAddr  string
	EtcdKey   string
	Collector []Collector
}

type Collector struct {
	LogPath  string `json:"logpath"`
	Topic    string `json:"topic"`
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

	// etcd config
	etcdAddr := conf.String("etcd::addr")
	if len(etcdAddr) == 0 {
		panic("etcdAddr is not defined!")
	}
	AppConfig.EtcdAddr = etcdAddr

	// etcd config
	configKey := conf.String("etcd::configKey")
	if len(configKey) == 0 {
		panic("configKey is not defined!")
	}
	AppConfig.EtcdKey = configKey

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
