package v1

import (
	"fmt"
	"net/http"

	"github.com/DesistDaydream/dtcg/pkg/dtcg/api/v1/models"
	deckprice "github.com/DesistDaydream/dtcg/pkg/dtcg/deck_price"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func PostDeckPrice(c *gin.Context) {
	// 允许跨域
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	var req models.PostDeckPriceRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ReqBodyErrorReponse{
			Message: "解析请求体异常",
			Data:    fmt.Sprintf("%v", err),
		})
		return
	}

	resp, err := deckprice.GetResp(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ReqBodyErrorReponse{
			Message: "获取响应失败",
			Data:    fmt.Sprintf("%v", err),
		})
		return
	}

	logrus.WithFields(logrus.Fields{
		"最低价": resp.MinPrice,
		"集换价": resp.AvgPrice,
	}).Debugf("卡组价格")

	c.JSON(200, &resp)
}

// 根据 card_id_from_db 获取卡组价格
func PostDeckPriceWithID(c *gin.Context) {
	// 允许跨域
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	var req models.PostDeckPriceWithIDReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ReqBodyErrorReponse{
			Message: "解析请求体异常",
			Data:    fmt.Sprintf("%v", err),
		})
		return
	}

	resp, err := deckprice.GetRespWithID(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ReqBodyErrorReponse{
			Message: "获取响应失败",
			Data:    fmt.Sprintf("%v", err),
		})
		return
	}

	logrus.WithFields(logrus.Fields{
		"最低价": resp.MinPrice,
		"集换价": resp.AvgPrice,
	}).Debugf("卡组价格")

	c.JSON(200, &resp)
}
