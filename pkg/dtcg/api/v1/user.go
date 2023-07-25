package v1

import (
	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetUser(c *gin.Context) {
	// 允许跨域
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	uid := c.Param("uid")

	resp, err := database.GetUser(uid)
	if err != nil {
		logrus.Errorf("获取用户信息失败，原因: %v", err)
	}

	c.JSON(200, &resp)
}
