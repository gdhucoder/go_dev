package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	fmt.Println("connect succ")
	defer cli.Close()

	cli.Put(context.Background(), "/logagent/conf/", "8888888")
	for {
		rch := cli.Watch(context.Background(), "/logagent/conf/")
		for wresp := range rch {
			for _, ev := range wresp.Events {
				fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}
		}
	}
	// PUT "/logagent/conf/" : "1"
	// PUT "/logagent/conf/" : "2"
	// PUT "/logagent/conf/" : "3"
	// PUT "/logagent/conf/" : "4"
	// PUT "/logagent/conf/" : "5"
	// PUT "/logagent/conf/" : "6"
	// PUT "/logagent/conf/" : "7"
	// PUT "/logagent/conf/" : "8"
	// PUT "/logagent/conf/" : "9"
}
