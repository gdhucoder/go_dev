package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go_dev/day11/logagent/conf"
	"time"

	"github.com/astaxie/beego/logs"
	etcd_client "go.etcd.io/etcd/clientv3"
)

type EtcdClient struct {
	client *etcd_client.Client
}

var (
	etcdClient *EtcdClient
)

// save config: logpath topic
var collectConf []conf.Collector

// move config setting in config file to ETCD
func InitEctd(ectdaddr string, key string) (collectConf []conf.Collector, err error) {

	cli, err := etcd_client.New(etcd_client.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	etcdClient = &EtcdClient{
		client: cli,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	for _, ip := range localIPArrayArr {
		etcdKey := fmt.Sprintf("%s%s", key, ip)
		logs.Debug("ectdKey:", etcdKey)
		resp, err := cli.Get(ctx, etcdKey)
		if err != nil {
			logs.Error("client get from etcd failed, err :%v", err)
			continue
		}
		logs.Debug("resp from etcd:%v", resp.Kvs)
		for _, v := range resp.Kvs {
			if string(v.Key) == etcdKey {
				err = json.Unmarshal(v.Value, &collectConf)
				if err != nil {
					logs.Error("unmarshal failed, err:", err)
					continue
				}
			}
			logs.Debug("log config is %v", collectConf)
		}

	}
	fmt.Printf("connect succ %v\n", cli)
	return
}
