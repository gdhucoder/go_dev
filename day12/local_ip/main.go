package main

import (
	"fmt"
	"net"
)

var (
	localIPArray []string
)

func init() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(fmt.Sprintf("get local ip failed, %v", err))
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				localIPArray = append(localIPArray, ipnet.IP.String())
			}
		}
	}

	fmt.Println(localIPArray)
	// 192.168.31.224
	// [169.254.34.53 169.254.252.17 172.17.183.1 169.254.122.129 172.26.66.1 192.168.31.224]
}

func main() {
	fmt.Println("init")
}
