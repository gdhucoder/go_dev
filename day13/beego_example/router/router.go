package router

import (
	"github.com/astaxie/beego"
	// D:\project\src\go_dev\day13\beego_example\controller\IndexController.go
	index "go_dev/day13/beego_example/controller"
)

func init() {
	beego.Router("/index", &index.IndexController{}, "*:Index")
	beego.Router("/api/index", &index.IndexController{}, "*:APIJson")
}
