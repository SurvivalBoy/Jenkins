package main

import (
	"github.com/Jenkins/router"
	"github.com/Jenkins/setting"
	"github.com/phjt-go/logger"
)

func main() {

	// 加载日志配置
	logger.SetLogger(setting.GetString("logger_jsonFile"))

	// 启动路由
	router.Run()
}
