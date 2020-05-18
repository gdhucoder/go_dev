package main

import (
	_ "go_dev/day14/SecKill/SecProxy/router"

	"github.com/astaxie/beego"
)

func main() {
	// 编译时会根据编译生成的路径寻找conf
	// 需要在SecProxy文件夹下编译
	// second skill inter
	err := initConfig()
	if err != nil {
		panic(err)
		return
	}

	err = initSec()
	if err != nil {
		panic(err)
	}
	beego.Run()
}
