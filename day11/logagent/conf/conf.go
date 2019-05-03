package conf

import (
	"fmt"

	"github.com/astaxie/beego/config"
)

type Config struct {
	LogLevel  string
	LogPath   string
	collector []Collector
}

type Collector struct {
	LogPath string
	Topic   string
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
	loadCollector(conf)
	return
}

func loadCollector(conf config.Configer) (err error) {
	collcetor := Collector{}
	collcetor.LogPath = conf.String("collect::log_path")
	collcetor.Topic = conf.String("collect::topic")
	AppConfig.collector = append(AppConfig.collector, collcetor)
	return
}
