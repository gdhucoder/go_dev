package main

import (
	"fmt"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
)

var (
	secKillConf = &SecSkillConf{}
)

type RedisConfig struct {
	redis_addr         string
	redis_max_idle     int
	redis_max_actlive  int
	redis_idle_timeout int
}

type SecSkillConf struct {
	redisConf RedisConfig
	etcdAddr  string
}

func initConfig() (err error) {
	redisAddr := beego.AppConfig.String("redis_addr")
	ectdAddr := beego.AppConfig.String("etcd_addr")
	logs.Debug("redis addr %v, etcd addr %v", redisAddr, ectdAddr)
	if len(redisAddr) == 0 || len(ectdAddr) == 0 {
		err = fmt.Errorf("init config failed, redis[%s] or ectd[%s] is nil",
			redisAddr, ectdAddr)
		return
	}
	secKillConf.redisConf.redis_addr = redisAddr
	secKillConf.etcdAddr = ectdAddr
	return
}
