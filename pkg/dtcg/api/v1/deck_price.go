package v1

import (
	"fmt"
	"net/http"

	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/pkg/dtcg/api/v1/models"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/services/community"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	m2 "github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/services/community/models"
)

// type fanxing interface {
// 	m2.Egg | m2.Main
// }

// func genData[T fanxing](resp models.PostDeckPriceResponse, card T) {

// }

func genData(card *m2.CardInfo, resp *models.PostDeckPriceResponse, c *gin.Context) {
	cardPrice, err := database.GetCardPrice(fmt.Sprint(card.Cards.CardID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.ReqBodyErrorReponse{
			Message: "获取卡片价格失败",
			Data:    "",
		})
		return
	}

	minPrice := cardPrice.MinPrice * float64(card.Number)
	avgPrice := cardPrice.AvgPrice * float64(card.Number)

	resp.Data = append(resp.Data, models.MutCardPrice{
		Count:          int(card.Number),
		Serial:         cardPrice.Serial,
		ScName:         cardPrice.ScName,
		AlternativeArt: cardPrice.AlternativeArt,
		MinPrice:       minPrice,
		AvgPrice:       avgPrice,
	})

	resp.MinPrice = resp.MinPrice + minPrice
	resp.AvgPrice = resp.AvgPrice + avgPrice
}

func PostDeckPrice(c *gin.Context) {
	// 允许跨域
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	var (
		req  models.PostDeckPriceRequest
		resp models.PostDeckPriceResponse
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ReqBodyErrorReponse{
			Message: "请求体错误",
			Data:    "",
		})
		return
	}

	client := community.NewCommunityClient(core.NewClient(""))
	decks, err := client.PostDeckConvert(req.Deck)
	if err != nil {
		logrus.Errorln(err)
		return
	}

	// TODO: 假设 Eggs 和 Main 最后是两种类型的话，怎么用泛型？
	for _, card := range decks.Data.DeckInfo.Eggs {
		genData(&card, &resp, c)
	}

	// TODO: 假设 Eggs 和 Main 最后是两种类型的话，怎么用泛型？
	for _, card := range decks.Data.DeckInfo.Main {
		genData(&card, &resp, c)
	}

	c.JSON(200, &resp)

	logrus.WithFields(logrus.Fields{
		"最低价": resp.MinPrice,
		"集换价": resp.AvgPrice,
	}).Infof("卡组价格")
}
