package http

import (
	"ExcelTools/reader"
	"ExcelTools/tools/settings"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

//	启动http监听服务
func Run(httpConfig *settings.HttpConfig) error {

	router := gin.Default()
	router.POST("/", handle)
	return router.Run(fmt.Sprintf(":%v", httpConfig.Port))
}
func handle(c *gin.Context) {
	type response struct {
		ErrorCode int    `json:"error_code"`
		Msg       string `json:"msg"`
	}
	req := reader.Request{}
	var b1 []byte
	_, err := c.Request.Body.Read(b1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b1))
	e := c.ShouldBind(&req)
	if settings.Conf.Mode != "prod" {
		bytes, e := json.Marshal(req)
		zap.L().Debug("请求数据", zap.ByteString("request", bytes), zap.Error(e))
	}
	if e != nil {
		zap.L().Error("接口数据格式错误", zap.Error(e))
		c.JSON(http.StatusBadRequest, &response{ErrorCode: 1, Msg: e.Error()})
		return
	}
	//	投递数据
	reader.HttpData <- &req
	c.JSON(http.StatusOK, &response{ErrorCode: 0})
}
