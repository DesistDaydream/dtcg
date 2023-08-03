package v1

import (
	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/pkg/dtcg/api/v1/models"
	"github.com/DesistDaydream/dtcg/pkg/dtcg/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Login(c *gin.Context) {
	// 获取客户端登录时的 IP，记录下这个数据，可以用于很多地方，比如根据 IP 判断登录失败次数达到阈值后禁止登录。
	ip := c.ClientIP()
	logrus.Infof("登录 IP: %v", ip)

	var req models.LoginReqBody
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorWithDataResp(c, err, 400, nil)
		return
	}

	// 根据用户名从数据库获取用户信息
	user, err := database.GetUserByName(req.Username)
	if err != nil {
		utils.ErrorWithDataResp(c, err, 400, nil)
		return
	}

	// 验证密码是否正确
	err = user.ValidatePassword(req.Password)
	if err != nil {
		utils.ErrorWithDataResp(c, err, 400, nil)
		return
	}

	// 生成 TOKEN
	token, err := utils.GenerateToken(user.Username)
	if err != nil {
		utils.ErrorWithDataResp(c, err, 400, nil, true)
		return
	}

	utils.SuccessResp(c, gin.H{"token": token})
}
