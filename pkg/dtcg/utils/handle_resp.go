package utils

import (
	"github.com/DesistDaydream/dtcg/pkg/dtcg/flags"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 一些通用的成功或失败的响应体
type Resp[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func ErrorWithDataResp(c *gin.Context, err error, code int, data interface{}, isStdout ...bool) {
	// 是否同步在服务端后台输出错误信息
	if len(isStdout) > 0 && isStdout[0] {
		if flags.F.Debug || flags.F.Dev {
			logrus.Errorf("%+v", err)
		} else {
			logrus.Errorf("%v", err)
		}
	}

	// 错误信息格式
	c.JSON(200, Resp[interface{}]{
		Code:    code,
		Message: err.Error(),
		Data:    data,
	})

	c.Abort()
}

func SuccessResp(c *gin.Context, data ...interface{}) {
	if len(data) == 0 {
		c.JSON(200, Resp[interface{}]{
			Code:    200,
			Message: "success",
			Data:    nil,
		})
		return
	}
	c.JSON(200, Resp[interface{}]{
		Code:    200,
		Message: "success",
		Data:    data[0],
	})
}
