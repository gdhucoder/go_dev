package main

import (
	"context"
	"fmt"
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

func InitEctd(ectdaddr string, key string) (err error) {

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
		for k, v := range resp.Kvs {
			fmt.Println("=============", k, v)
		}

	}
	fmt.Printf("connect succ %v\n", cli)
	return
}
