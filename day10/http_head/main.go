package main

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

var url = []string{
	"http://www.baidu.com",
	"http://google.com",
	"http://taobao.com",
}

func main() {

	c := http.Client{
		Transport: &http.Transport{
			Dial: func(network, addr string) (net.Conn, error) {
				timeout := time.Second * 2
				return net.DialTimeout(network, addr, timeout)
			},
		},
	}

	for _, v := range url {
		resp, err := c.Head(v)
		if err != nil {
			fmt.Printf("head %s failed, err:%v\n", v, err)
			continue
		}

		// head succ, status:200 OK
		// head http://google.com failed, err:Head http://google.com: dial tcp 216.58.200.14:80: i/o timeout
		// head succ, status:200 OK

		fmt.Printf("head succ, status:%v\n", resp.Status)
	}
}
