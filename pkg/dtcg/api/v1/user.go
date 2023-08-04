package v1

import (
	"fmt"
	"net/http"

	"github.com/DesistDaydream/dtcg/config"
	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/pkg/dtcg/api/v1/models"
	"github.com/DesistDaydream/dtcg/pkg/dtcg/utils"
	"github.com/DesistDaydream/dtcg/pkg/handler"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ListUser(c *gin.Context) {
	var reqQuery models.CommonReqQuery
	if err := c.ShouldBindQuery(&reqQuery); err != nil {
		logrus.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if reqQuery.PageSize == 0 || reqQuery.PageNum == 0 {
		reqQuery.PageSize = 10
		reqQuery.PageNum = 1
	}

	resp, err := database.ListUser(reqQuery.PageSize, reqQuery.PageNum)
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

// 设置当前登录的用户信息
func CurrentUser(c *gin.Context) {
	token := c.GetHeader("Authorization")
	logrus.Debugf("检查传入的 TOKEN: %v", token)

	if token == "" {
		utils.ErrorWithDataResp(c, fmt.Errorf("token is empty"), 401, nil, true)
		return
	}

	userClaims, err := utils.ParseToken(token)
	if err != nil {
		utils.ErrorWithDataResp(c, err, 401, nil)
		return
	}

	user, err := database.GetUserByName(userClaims.Username)
	if err != nil {
		utils.ErrorWithDataResp(c, err, 401, nil)
		return
	}

	// 实例化处理器以便后续代码可以从 卡查 和 集换社 中获取信息
	handler.H = handler.NewHandler(user, true, config.Conf.Moecard.Retry)

	// 设定当前用户信息以便在其他部分代码中获取
	c.Set("user", user)

	c.Next()
}
