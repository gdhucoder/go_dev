package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	etcd_client "go.etcd.io/etcd/clientv3"
)

const (
	EtcdKey = "/logagent/config/192.168.31.224"
)

type LogConf struct {
	Path  string `json:"path"`
	Topic string `json:"topic"`
}

func SetLogConfigToEtcd() {
	cli, err := etcd_client.New(etcd_client.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	fmt.Println("connect succ")
	defer cli.Close()

	var longConfArr []LogConf

	longConfArr = append(
		longConfArr,
		LogConf{
			Path:  "D:/project/src/go_dev/day11/logagent/logs/logagent.log",
			Topic: "test",
		},
	)
	data, err := json.Marshal(longConfArr)
	if err != nil {
		panic("init config failed")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = cli.Put(ctx, EtcdKey, string(data))
	if err != nil {
		fmt.Println("etcd put failed, err:", err)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := cli.Get(ctx, EtcdKey)
	if err != nil {
		fmt.Println("etcd put failed, err:", err)
		return
	}

	for _, ev := range resp.Kvs {
		fmt.Printf("%s: %s", ev.Key, ev.Value)
	}

}

func main() {
	SetLogConfigToEtcd()
	// /logagent/config/: [{"path":"D:\\project\\src\\go_dev\\day11\\logagent\\logs\\logagent.log","topic":"test"}]
}
