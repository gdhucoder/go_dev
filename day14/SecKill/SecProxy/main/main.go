package main

import (
	"fmt"
	_ "go_dev/day14/SecKill/SecProxy/router"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
)

func initConfig() (err error) {
	redisAddr := beego.AppConfig.String("redis_addr")
	ectdAddr := beego.AppConfig.String("etcd_addr")
	logs.Debug("redis addr %v, etcd addr %v", redisAddr, ectdAddr)
	if len(redisAddr) == 0 || len(ectdAddr) == 0 {
		err = fmt.Errorf("init config failed, redis[%s] or ectd[%s] is nil",
			redisAddr, ectdAddr)
		return
	}
	secKillConf.redisAddr = redisAddr
	secKillConf.etcdAddr = ectdAddr
	return
}

func main() {
	// 编译时会根据编译生成的路径寻找conf
	// 需要在SecProxy文件夹下编译
	// second skill inter
	err := initConfig()
	if err != nil {
		panic(err)
		return
	}
	beego.Run()
}
