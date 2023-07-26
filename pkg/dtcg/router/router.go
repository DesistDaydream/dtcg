package router

import (
	v1 "github.com/DesistDaydream/dtcg/pkg/dtcg/api/v1"
	"github.com/DesistDaydream/dtcg/pkg/dtcg/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	api := r.Group("/api/v1")
	api.GET("/user/info/:uid", v1.GetUser)

	api.POST("/set/desc", v1.PostCardSets)

	api.GET("/card/desc", v1.GetCardsDesc)
	api.POST("/card/desc", v1.PostCardsDesc)

	api.GET("/card/price", v1.GetCardsPrice)
	api.POST("/card/price", v1.PostCardsPrice)
	api.POST("/card/pricewithimg", v1.PostCardsPriceWithDtcgDBImg)

	api.POST("/deck/price/json", v1.PostDeckPriceWithJSON)
	api.GET("/deck/price/hid/:hid", v1.GetDeckPriceWithHID)
	api.GET("/deck/price/cdid/:cdid", v1.GetDeckPriceWithCloudDeckID)

	// 将 DTCG DB 卡组广场中卡组的 URL 中最后的 HID，转变为由 card_id_from_db 组成的纯字符串格式。
	api.GET("/deck/converter/:hid", v1.GetDeckConverter)
	// 根据上面转换后的字符串格式的卡组信息，获取卡组价格。
	api.POST("/deck/price/ids", v1.PostDeckPriceWithIDS)

	api.POST("/login", v1.Login)

	auth := api.Group("", middlewares.Auth)
	auth.GET("/auth/test", v1.TestAuth)
	return r
}
