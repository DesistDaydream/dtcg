package v1

import (
	"net/http"

	dbmodels "github.com/DesistDaydream/dtcg/internal/database/models"

	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/pkg/dtcg/api/v1/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 获取所有卡牌的描述
func GetCardsDesc(c *gin.Context) {

	// 绑定 url query
	var req models.CommonReqQuery

	if err := c.ShouldBindQuery(&req); err != nil {
		logrus.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.PageSize == 0 || req.PageNum == 0 {
		req.PageSize = 10
		req.PageNum = 1
	}

	resp, err := database.GetCardsDesc(req.PageSize, req.PageNum)
	if err != nil {
		logrus.Errorf("%v", err)
	}

	c.JSON(200, resp)
}

// 根据条件获取卡牌描述详情
func PostCardsDesc(c *gin.Context) {
	// 绑定 url query
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

	// 绑定请求体
	var reqBody dbmodels.CardDescQuery
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		logrus.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ReqBodyErrorResp{
			Message: "请求体错误",
			Data:    "",
		})
		return
	}

	resp, err := database.GetCardDescByCondition(reqQuery.PageSize, reqQuery.PageNum, &reqBody)
	if err != nil {
		logrus.Errorf("%v", err)
	}

	c.JSON(200, resp)
}
