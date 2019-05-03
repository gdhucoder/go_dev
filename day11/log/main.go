package main

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego/logs"
)

func main() {
	config := make(map[string]interface{})
	config["filename"] = "D:\\project\\src\\go_dev\\day11\\log\\logcollect.log"
	config["level"] = logs.LevelDebug

	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("marshal failed, err:", err)
		return
	}

	logs.SetLogger(logs.AdapterFile, string(configStr)) // 转成json字符串

	logs.Debug("this is a test, my name is %s", "stu01")
	logs.Trace("this is a trace, my name is %s", "stu02")
	logs.Warn("this is a warn, my name is %s", "stu03")

	// 2019/05/02 22:39:25.952 [W]  this is a warn, my name is stu03
	// 2019/05/02 22:40:33.824 [W]  this is a warn, my name is stu03
	// 2019/05/02 22:41:38.942 [W]  this is a warn, my name is stu03
	// 2019/05/02 22:41:57.560 [D]  this is a test, my name is stu01
	// 2019/05/02 22:41:57.560 [D]  this is a trace, my name is stu02
	// 2019/05/02 22:41:57.560 [W]  this is a warn, my name is stu03

	// 可以当做错误日志。定一个错误的topic，每分钟检查一下，若出现错误发送邮件。

	// 热加载配置文件，实时检查文件的变更，程序不需要重启。
}
