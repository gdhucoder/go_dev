package main

import (
	"fmt"
	"time"

	etcd_client "go.etcd.io/etcd/clientv3"
)

func InitEctd() (err error) {

	cli, err := etcd_client.New(etcd_client.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	fmt.Printf("connect succ %v\n", cli)
	return
}
