package router

import (
	"go_dev/day14/SecKill/SecProxy/controller"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/seckill", &controller.SkillController{}, "*:SecKill")
	beego.Router("/secinfo", &controller.SkillController{}, "*:SecInfo")
}
