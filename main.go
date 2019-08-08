package main

import (
	"github.com/phjt-go/logger"
	"jenkins_demo/router"
	"jenkins_demo/setting"
)

func main() {

	// 加载日志配置
	logger.SetLogger(setting.GetString("logger_jsonFile"))

	// 启动路由
	router.Run()
}
