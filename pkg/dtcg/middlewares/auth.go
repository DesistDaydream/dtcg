package middlewares

import (
	"fmt"

	"github.com/DesistDaydream/dtcg/pkg/dtcg/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Auth(c *gin.Context) {
	token := c.GetHeader("Authorization")
	logrus.Infof("%v", token)

	if token == "" {
		utils.ErrorResp(c, fmt.Errorf("token 为空"), 401, true)
		return
	}

	c.Next()
}
