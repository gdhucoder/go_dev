package main

import (
	"fmt"
	"go_dev/day11/logagent/conf"
	"go_dev/day11/logagent/tailf"

	"github.com/astaxie/beego/logs"
)

func main() {
	// load config
	err := conf.InitConf("ini", "D:\\project\\src\\go_dev\\day11\\logagent\\conf\\logagent.conf")
	if err != nil {
		fmt.Printf("init config failed, err: %v\n", err)
	}
	fmt.Println(conf.AppConfig)

	// init logger
	err = InitLogger()
	if err != nil {
		fmt.Printf("init logger failed, err: %v\n", err)
	}
	logs.Debug("init logger success!")

	// init tailf
	err = tailf.InitTailF()
	if err != nil {
		fmt.Printf("init tailf failed, err: %v\n", err)
	}
	logs.Debug("init tailf success!")

}
