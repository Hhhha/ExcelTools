package server

import (
	"ExcelTools/reader/http"
	"ExcelTools/tools/settings"
	"go.uber.org/zap"
)

//	启动http服务
func startHttpServer(config *settings.HttpConfig) {
	if config.Start != true {
		zap.L().Debug("http服务未开启")
		return
	}
	if err := http.Run(config); err != nil {
		zap.L().Debug("http服务启动失败", zap.Error(err))
		return
	}
}
