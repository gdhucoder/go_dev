package main

import (
	"fmt"

	"github.com/astaxie/beego/config"
)

func main() {
	conf, err := config.NewConfig("ini", "D:\\project\\src\\go_dev\\day11\\config\\logagent.conf")
	if err != nil {
		fmt.Println("new config failed, err:", err)
		return
	}

	ip := conf.String("server::listen_ip")
	fmt.Println("ip:", ip)

	port, err := conf.Int("server::listen_port")
	if err != nil {
		fmt.Println("read server:port failed, err:", port)
		return
	}
	fmt.Println("port:", port)

	log_level, err := conf.Int("logs::log_level")
	if err != nil {
		fmt.Println("read log_level failed, ", err)
		return
	}
	fmt.Println("log_level:", log_level)

	log_path := conf.String("collect::log_path")
	fmt.Println("collect:", log_path)
	// ip: 0.0.0.0
	// port: 8080
	// log_level: 5
	// collect: src/go_dev/day11/config/logagent.conf
}
