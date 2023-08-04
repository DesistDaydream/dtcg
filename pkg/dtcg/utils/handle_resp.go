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

// 模拟 Gin 的 c.AbortWithXXX 方法的逻辑，当有错误是返回一个格式统一的信息。
// 主要是这种错误信息自主可控，避免 Gin 的一些不足；也算是向 Gin 的学习。
// 代码参考 Alist。
func ErrorWithDataResp(c *gin.Context, err error, code int, data interface{}, isStdout ...bool) {
	// 是否同步在服务端后台输出错误信息
	if len(isStdout) > 0 && isStdout[0] {
		if flags.F.Debug || flags.F.Dev {
			logrus.Errorf("%+v", err)
		} else {
			logrus.Errorf("%v", err)
		}
	}

	c.Abort()

	// 错误信息格式
	c.JSON(200, Resp[interface{}]{
		Code:    code,
		Message: err.Error(),
		Data:    data,
	})
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
