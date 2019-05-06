package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go_dev/day11/logagent/conf"
	"time"

	"go.etcd.io/etcd/mvcc/mvccpb"

	"github.com/astaxie/beego/logs"
	etcd_client "go.etcd.io/etcd/clientv3"
)

type EtcdClient struct {
	client *etcd_client.Client
	keys   []string
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

		// save keys in a slice
		etcdClient.keys = append(etcdClient.keys, etcdKey)

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

	initEtcdWatcher()
	return
}

func initEtcdWatcher() {
	for _, key := range etcdClient.keys {
		logs.Debug("ectd is watching: ", key)
		go watchKey(key)
	}
}

func watchKey(key string) {
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
	for {
		rch := cli.Watch(context.Background(), key)
		var changedCollectConf []conf.Collector
		var getConfSucc = true
		for wresp := range rch {
			for _, ev := range wresp.Events {
				if ev.Type == mvccpb.DELETE {
					logs.Warn("key[%s] is deleted", key)
					continue
				}
				// PUT "/logagent/config/192.168.31.224" : "[{\"logpath\":\"D:/project/src/go_dev/day11/logagent/logs/logagent.log\",\"topic\":\"test\",\"ChanSize\":0}]"
				if ev.Type == mvccpb.PUT && string(ev.Kv.Key) == key {
					err = json.Unmarshal(ev.Kv.Value, &changedCollectConf)
					if err != nil {
						logs.Error("key [%s], Unmarshal[%s], err:%v ", key, ev.Kv.Value, err)
						getConfSucc = false
						continue
					}
				}
				// Type:PUT,DELETE
				logs.Debug("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}

			// if true, update config
			if getConfSucc {
				tailf.UpdateConfig(changedCollectConf)
			}
		}
	}
}
