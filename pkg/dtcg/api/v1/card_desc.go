package v1

import (
	"net/http"

	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/pkg/dtcg/api/v1/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func PostCardsDesc(c *gin.Context) {
	// 允许跨域
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	var req models.PostCardsDescReq

	if err := c.ShouldBindJSON(&req); err != nil {
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
