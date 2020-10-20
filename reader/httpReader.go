package reader

import (
	"ExcelTools/tools/settings"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

//	启动http监听服务
func Run(httpConfig *settings.HttpConfig) error {

	router := gin.Default()
	router.POST("/", func(context *gin.Context) {

	})
	return router.Run(fmt.Sprintf(":%v", httpConfig.Port))
}
func handle(c *gin.Context) {
	type response struct {
		ErrorCode int    `json:"error_code"`
		Msg       string `json:"msg"`
	}
	req := &Request{}
	e := c.ShouldBind(req)
	if e != nil {
		zap.L().Error("接口数据格式错误", zap.Error(e))
		c.JSON(http.StatusBadRequest, &response{ErrorCode: 1, Msg: e.Error()})
	}
	c.JSON(http.StatusOK, &response{ErrorCode: 0})
}
