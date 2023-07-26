package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Resp[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type ReqBodyErrorResp struct {
	Message string `json:"message"`
	Data    string `json:"data"`
}

func ErrorResp(c *gin.Context, err error, code int, l ...bool) {
	ErrorWithDataResp(c, err, code, nil, l...)
}

func ErrorWithDataResp(c *gin.Context, err error, code int, data interface{}, l ...bool) {
	if len(l) > 0 && l[0] {
		// if flags.Debug || flags.Dev {
		// 	logrus.Errorf("%+v", err)
		// } else {
		// 	logrus.Errorf("%v", err)
		// }
		logrus.Errorf("%v", err)
	}
	c.JSON(200, Resp[interface{}]{
		Code:    code,
		Message: err.Error(),
		Data:    data,
	})
	c.Abort()
}
