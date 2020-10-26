package main

import (
	"ExcelTools/server"
	"ExcelTools/tools/logger"
	"ExcelTools/tools/settings"
	"encoding/json"
	"go.uber.org/zap"
)

func main() {
	e := settings.Init()
	if e != nil {
		zap.L().Fatal("配置文件初始化失败", zap.Error(e))
		return
	}
	bytes, _ := json.Marshal(settings.Conf)
	zap.L().Debug("配置文件初始化", zap.ByteString("settings", bytes))
	e = logger.Init(settings.Conf.LogConfig, settings.Conf.Mode)
	if e != nil {
		zap.L().Fatal("日志模块初始化失败", zap.Error(e))
		return
	}
	zap.L().Debug("日志模块初始化成功")

	//	启动服务
	server.Start()
	select {}
}
