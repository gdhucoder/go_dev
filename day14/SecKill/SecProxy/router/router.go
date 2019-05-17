package router

import (
	"go_dev/day14/SecKill/SecProxy/controller"

	"github.com/astaxie/beego"
)

func init() {
	// 秒杀 second kill
	beego.Router("/seckill", &controller.SkillController{}, "*:SecKill")
	// 秒杀信息 info
	beego.Router("/secinfo", &controller.SkillController{}, "*:SecInfo")
}
