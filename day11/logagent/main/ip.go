package main

import (
	"fmt"
	"net"
)

var (
	localIPArrayArr []string
)

func init() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(fmt.Sprintf("get local ip failed, %v", err))
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				localIPArrayArr = append(localIPArrayArr, ipnet.IP.String())
			}
		}
	}

	fmt.Println(localIPArrayArr)
}
