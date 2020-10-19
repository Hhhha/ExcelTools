package main

import (
	"ExcelTools/tools/logger"
	"ExcelTools/tools/settings"
	"fmt"
	"go.uber.org/zap"
)

func main() {
	settings.Init()
	fmt.Printf("========================%v", settings.Conf.LogConfig)
	logger.Init(settings.Conf.LogConfig, settings.Conf.Mode)
	zap.L().Info("===============================131")
}
