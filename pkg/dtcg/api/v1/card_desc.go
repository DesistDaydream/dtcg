package v1

import (
	"net/http"

	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/pkg/dtcg/api/v1/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetCardsDesc(c *gin.Context) {
	// 允许跨域
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	// 绑定 url query
	var req models.GetCardsDescReqQuery

	if err := c.ShouldBindQuery(&req); err != nil {
		logrus.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.PageSize == 0 || req.PageNum == 0 {
		req.PageSize = 10
		req.PageNum = 1
	}

	resp, err := database.GetCardDesc(req.PageSize, req.PageNum)
	if err != nil {
		logrus.Errorf("%v", err)
	}

	c.JSON(200, resp)
}

func PostCardsDesc(c *gin.Context) {
	// 允许跨域
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	var req models.PostCardsDescReq

	if err := c.ShouldBindJSON(&req); err != nil {
		logrus.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ReqBodyErrorResp{
			Message: "请求体错误",
			Data:    "",
		})
		return
	}

	resp, err := database.GetCardDesc(req.PageSize, req.PageNum)
	if err != nil {
		logrus.Errorf("%v", err)
	}

	c.JSON(200, resp)
}
