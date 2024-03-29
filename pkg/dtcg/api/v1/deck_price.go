package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/pkg/dtcg/api/v1/models"
	deckprice "github.com/DesistDaydream/dtcg/pkg/dtcg/deck_price"
	"github.com/DesistDaydream/dtcg/pkg/dtcg/utils"
	"github.com/DesistDaydream/dtcg/pkg/handler"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 根据 DTCG_DB 导出的 JSON 格式卡组信息获取卡组价格
func PostDeckPriceWithJSON(c *gin.Context) {
	// 绑定请求体
	var req models.PostDeckPriceWithJSONReqBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ReqBodyErrorResp{
			Message: "解析请求体异常",
			Data:    fmt.Sprintf("%v", err),
		})
		return
	}

	resp, err := deckprice.GetDeckPriceWithJSON(&req)
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
	hid := c.Param("hid")

	decks, err := handler.H.MoecardServices.Community.GetDeck(hid)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ReqBodyErrorResp{
			Message: "无法获取到卡组信息",
			Data:    fmt.Sprintf("%v", err),
		})
		return
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

	resp, err := deckprice.GenDeckPriceWithMoecardID(&req)
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

// 根据云 Cloud Deck ID(云卡组ID) 获取卡组价格，云卡组ID是个人页面的卡组ID，必须携带登录 Token 才可以获取到
// 这种获取方式是最完整的，但是也是很麻烦的，因为需要登录，而且只能是自己的卡组。
func GetDeckPriceWithCloudDeckID(c *gin.Context) {
	cloudDeckID := c.Param("cdid")

	decks, err := handler.H.MoecardServices.Community.GetDeckCloud(cloudDeckID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ReqBodyErrorResp{
			Message: "无法获取到云卡组信息",
			Data:    fmt.Sprintf("%v", err),
		})
		return
	}

	if len(decks.Data.DeckInfo.Eggs) == 0 && len(decks.Data.DeckInfo.Main) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ReqBodyErrorResp{
			Message: "卡组为空，请添加至少一张卡牌",
			Data:    fmt.Sprintf("%v", err),
		})
		return
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

	resp, err := deckprice.GenDeckPriceWithMoecardID(&req)
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

// 根据集换社的心愿单 ID 获取卡组价格，可以通过分享心愿单中的风险链接找到心愿单 ID
func GetDeckPriceWithJHSWishListID(c *gin.Context) {
	wishListID := c.Param("wlid")
	var (
		cardsID []string
		page    int = 1
	)

	for {
		wishListGetResp, err := handler.H.JhsServices.Wishes.Get(wishListID, page)
		if err != nil {
			logrus.Errorf("获取我想收清单失败，原因: %v", err)
			utils.ErrorWithDataResp(c, err, 400, nil)
			return
		}

		for _, card := range wishListGetResp.Data {
			cardPrice, err := database.GetCardPriceWhereCardVersionID(strconv.Itoa(card.CardVersionID))
			if err != nil {
				logrus.Errorf("%v", err)
			}
			for i := 0; i < card.Quantity; i++ {
				cardsID = append(cardsID, fmt.Sprint(cardPrice.CardIDFromDB))
			}
		}

		if wishListGetResp.NextPageURL == "" {
			logrus.Debugf("最后页: %v; 当前页: %v", wishListGetResp.LastPage, wishListGetResp.CurrentPage)
			break
		}

		// 每处理完一页，下一个循环需要处理的页+1
		page++
	}

	cardsIDString, _ := json.Marshal(&cardsID)

	req := models.PostDeckPriceWithIDReq{
		IDs: string(cardsIDString),
	}

	resp, err := deckprice.GenDeckPriceWithMoecardID(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ReqBodyErrorResp{
			Message: "获取响应失败",
			Data:    fmt.Sprintf("%v", err),
		})
		utils.ErrorWithDataResp(c, fmt.Errorf("由于 %v 原因导致根据 ID 计算", err), 400, nil)
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
	// 绑定请求体
	var req models.PostDeckPriceWithIDReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ReqBodyErrorResp{
			Message: "解析请求体异常",
			Data:    fmt.Sprintf("%v", err),
		})
		return
	}

	resp, err := deckprice.GenDeckPriceWithMoecardID(&req)
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

// 根据分享卡组的 ID 获取卡组价格，分享卡组的 ID 是网页端卡组编辑页面中，点击分享后获得的。
func GetDeckPriceWithShareID(c *gin.Context) {
	shareDeckID := c.Param("shareid")

	decks, err := handler.H.MoecardServices.Community.GetShareDeck(shareDeckID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ReqBodyErrorResp{
			Message: "无法获取到云卡组信息",
			Data:    fmt.Sprintf("%v", err),
		})
		return
	}

	if len(decks.Data.DeckInfo.Eggs) == 0 && len(decks.Data.DeckInfo.Main) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ReqBodyErrorResp{
			Message: "卡组为空，请添加至少一张卡牌",
			Data:    fmt.Sprintf("%v", err),
		})
		return
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

	resp, err := deckprice.GenDeckPriceWithMoecardID(&req)
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
