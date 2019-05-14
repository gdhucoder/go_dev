package controller

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type IndexController struct {
	beego.Controller
}

func (p *IndexController) Index() {
	logs.Debug("enter index controller")
	p.TplName = "index/index.html"
}

func (p *IndexController) APIJson() {
	logs.Debug("enter APIJson controller")
	m := make(map[string]interface{})
	m["code"] = 200
	m["name"] = "huguodong"
	p.Data["json"] = m
	p.ServeJSON(true)
}
