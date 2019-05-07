package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go_dev/day11/logagent/conf"
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

// tools : set config to etcd
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

	var longConfArr []conf.Collector

	// we can collect different levles of log (logpath, and topic)
	longConfArr = append(
		longConfArr,
		conf.Collector{
			LogPath: "D:/project/src/go_dev/day11/logagent/logs/logagent.log",
			Topic:   "test",
		},
	)

	longConfArr = append(
		longConfArr,
		conf.Collector{
			LogPath: "D:/project/src/go_dev/day11/logagent/logs/logagent_error.log",
			Topic:   "test_error",
		},
	)

	data, err := json.Marshal(longConfArr)
	if err != nil {
		panic("init config failed")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// cli.Delete(ctx, EtcdKey)
	// fmt.Println("delete success!")
	// return

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
