package server

import (
	"ExcelTools/tools/settings"
)

//	启动对应服务
func Start() {
	go startHttpServer(settings.Conf.HttpConfig)
}
