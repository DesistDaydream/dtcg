package v1

import (
	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/pkg/dtcg/api/v1/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 根据 User ID 获取指定用户的信息
func GetUser(c *gin.Context) {
	// 允许跨域
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	uid := c.Param("uid")

	userInfo, err := database.GetUser(uid)
	if err != nil {
		logrus.Errorf("获取用户信息失败，原因: %v", err)
	}

	c.JSON(200, &models.UserData{
		ID:           userInfo.ID,
		Username:     userInfo.Username,
		MoecardToken: userInfo.MoecardToken,
		JhsToken:     userInfo.JhsToken,
		CreatedAt:    userInfo.UpdatedAt,
		UpdatedAt:    userInfo.UpdatedAt,
	})
}
