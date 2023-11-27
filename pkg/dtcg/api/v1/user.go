package v1

import (
	"fmt"
	"strconv"

	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/pkg/dtcg/api/v1/models"
	"github.com/DesistDaydream/dtcg/pkg/dtcg/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ListUser(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	pageNum, _ := strconv.Atoi(c.DefaultQuery("page_num", "1"))

	resp, err := database.ListUser(pageSize, pageNum)
	if err != nil {
		logrus.Errorf("列出用户信息失败，原因: %v", err)
		utils.ErrorWithDataResp(c, fmt.Errorf("列出用户信息失败，原因: %v", err), 400, nil, true)
	}

	c.JSON(200, resp)
}

// 根据 User ID 获取指定用户的信息
func GetUser(c *gin.Context) {
	uid := c.Param("uid")

	userInfo, err := database.GetUser(uid)
	if err != nil {
		logrus.Errorf("获取用户信息失败，原因: %v", err)
		utils.ErrorWithDataResp(c, fmt.Errorf("获取用户信息失败，原因: %v", err), 400, nil, true)
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
