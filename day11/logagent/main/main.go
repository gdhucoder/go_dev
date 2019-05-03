package main

import (
	"fmt"
	"go_dev/day11/logagent/conf"
	"go_dev/day11/logagent/kafka"
	"go_dev/day11/logagent/tailf"

	"github.com/astaxie/beego/logs"
)

func main() {
	// load config
	err := conf.InitConf("ini", "D:\\project\\src\\go_dev\\day11\\logagent\\conf\\logagent.conf")
	if err != nil {
		fmt.Printf("init config failed, err: %v\n", err)
		return
	}

	// init logger
	err = InitLogger()
	if err != nil {
		fmt.Printf("init logger failed, err: %v\n", err)
		return
	}
	logs.Debug("init logger success!")

	// init tailf
	err = tailf.InitTailF(conf.AppConfig.Collector)
	if err != nil {
		fmt.Printf("init tailf failed, err: %v\n", err)
		return
	}

	err = kafka.InitKafka(conf.AppConfig.KafkaAddr)
	if err != nil {
		fmt.Printf("init kafka failed, err: %v\n", err)
		return
	}

	// go func() {
	// 	var count int
	// 	for {
	// 		count++
	// 		logs.Debug("test msg ", count)
	// 		time.Sleep(time.Second)
	// 	}
	// }()

	// &{debug D:\project\src\go_dev\day11\logagent\logs\logagent.log [{D:\project\src\go_dev\day11\logagent\logs\collect nginx_log}]}

	runServer()

}
