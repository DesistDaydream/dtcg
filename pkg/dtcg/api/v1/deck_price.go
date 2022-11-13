package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DesistDaydream/dtcg/pkg/dtcg/api/v1/models"
	deckprice "github.com/DesistDaydream/dtcg/pkg/dtcg/deck_price"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/services/community"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 根据 DTCG_DB 导出的 JSON 格式卡组信息获取卡组价格
func PostDeckPriceWithJSON(c *gin.Context) {
	// 允许跨域
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	var req models.PostDeckPriceWithJSONReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ReqBodyErrorResp{
			Message: "解析请求体异常",
			Data:    fmt.Sprintf("%v", err),
		})
		return
	}

	resp, err := deckprice.GetResp(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ReqBodyErrorResp{
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

// 根据所有卡牌的 card_id_from_db 获取卡组价格。这里的 card_id_from_db 是通过 GetDeckConverter() 函数获取的，也就是 /deck/converter/:hid 接口
func PostDeckPriceWithIDS(c *gin.Context) {
	// 允许跨域
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	var req models.PostDeckPriceWithIDReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ReqBodyErrorResp{
			Message: "解析请求体异常",
			Data:    fmt.Sprintf("%v", err),
		})
		return
	}

	resp, err := deckprice.GetRespWithID(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ReqBodyErrorResp{
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

// 根据 HID 获取卡组价格
func GetDeckPriceWithHID(c *gin.Context) {
	// 允许跨域
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	hid := c.Param("hid")

	client := community.NewCommunityClient(core.NewClient(""))
	decks, err := client.GetDeck(hid)
	if err != nil {
		logrus.Errorln(err)
	}

	var cardsID []string

	for _, card := range decks.Data.DeckInfo.Eggs {
		for i := 0; i < card.Number; i++ {
			cardsID = append(cardsID, fmt.Sprint(card.Cards.CardID))
		}
	}

	for _, card := range decks.Data.DeckInfo.Main {
		for i := 0; i < card.Number; i++ {
			cardsID = append(cardsID, fmt.Sprint(card.Cards.CardID))
		}
	}

	cardsIDString, _ := json.Marshal(&cardsID)

	req := models.PostDeckPriceWithIDReq{
		IDs: string(cardsIDString),
	}

	resp, err := deckprice.GetRespWithID(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ReqBodyErrorResp{
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
