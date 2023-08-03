package middlewares

import (
	"fmt"

	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/pkg/dtcg/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Auth(c *gin.Context) {
	token := c.GetHeader("Authorization")
	logrus.Debugf("检查传入的 TOKEN: %v", token)

	if token == "" {
		utils.ErrorWithDataResp(c, fmt.Errorf("token is empty"), 401, nil, true)
		return
	}

	userClaims, err := utils.ParseToken(token)
	if err != nil {
		utils.ErrorWithDataResp(c, err, 401, nil)
		c.Abort()
		return
	}

	user, err := database.GetUserByName(userClaims.Username)
	if err != nil {
		utils.ErrorWithDataResp(c, err, 401, nil)
		c.Abort()
		return
	}

	// 设定当前用户信息以便在其他部分代码中获取
	c.Set("user", user)

	logrus.Infof("当前登录的用户为: %v", user.Username)

	c.Next()
}
