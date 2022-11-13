package router

import (
	v1 "github.com/DesistDaydream/dtcg/pkg/dtcg/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	rr := r.Group("/api/v1")
	rr.POST("/set/desc", v1.PostCardSets)
	rr.POST("/card/desc", v1.PostCardsDesc)
	rr.POST("/card/price", v1.PostCardsPrice)
	rr.POST("/deck/price/json", v1.PostDeckPriceWithJSON)
	rr.GET("/deck/price/hid/:hid", v1.GetDeckPriceWithHID)

	// 将 DTCG DB 卡组广场中卡组的 URL 中最后的 HID，转变为由 card_id_from_db 组成的纯字符串格式。
	rr.GET("/deck/converter/:hid", v1.GetDeckConverter)
	// 根据上面转换后的字符串格式的卡组信息，获取卡组价格。
	rr.POST("/deck/price/ids", v1.PostDeckPriceWithIDS)

	return r
}
