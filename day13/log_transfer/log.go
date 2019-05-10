package main

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego/logs"
)

func initLogger() (err error) {
	config := make(map[string]interface{})
	config["filename"] = logConfig.LogPath
	config["level"] = convertLogLevel(logConfig.LogLevel)
	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("marshal failed, err:", err)
		return
	}
	logs.SetLogger(logs.AdapterFile, string(configStr)) // 转成json字符串
	return
}

func convertLogLevel(level string) int {
	switch level {
	case "debug":
		return logs.LevelDebug
	case "trace":
		return logs.LevelTrace
	case "info":
		return logs.LevelInfo
	case "warn":
		return logs.LevelWarn
	}
	return logs.LevelDebug
}
