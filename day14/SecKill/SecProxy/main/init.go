package main

import (
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/gomodule/redigo/redis"
)

var (
	redisPool *redis.Pool
)

func initRedis() (err error) {
	redisPool = &redis.Pool{
		MaxIdle:     secKillConf.redisConf.redis_max_idle,
		MaxActive:   secKillConf.redisConf.redis_max_actlive,
		IdleTimeout: time.Duration(secKillConf.redisConf.redis_idle_timeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", secKillConf.redisConf.redis_addr)
		},
	}

	conn := redisPool.Get()
	defer conn.Close()

	_, err = conn.Do("ping")
	if err != nil {
		logs.Error("ping redis failed, err:%v", err)
		return
	}

	
	// sex, _ := redis.String(conn.Do("GET", "huguodong"))
	// logs.Info("sex %s", sex)

	return
}

func initEtcd() (err error) {
	return
}

func initSec() (err error) {
	err = initRedis()
	logs.Info("%v", redisPool)
	if err != nil {
		logs.Error("init redis failed, err: %v", err)
		return
	}

	err = initEtcd()
	if err != nil {
		logs.Error("init etcd failed, err: %v", err)
		return
	}
	logs.Info("init second kill succ")
	return
}
