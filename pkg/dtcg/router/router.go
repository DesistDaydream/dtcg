package router

import (
	v1 "github.com/DesistDaydream/dtcg/pkg/dtcg/api/v1"
	"github.com/DesistDaydream/dtcg/pkg/dtcg/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func CommonSet(c *gin.Context) {
	method := c.Request.Method
	logrus.Debugf("设置一些通用头，检查请求方法: %v", method)
	// 允许跨域
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	// c.Header("Access-Control-Expose-Headers", "*")
	// c.Header("Access-Control-Allow-Credentials", "false")
	// c.Header("content-type", "application/json")

	c.Next()
}

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Any("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// API
	api := r.Group("/api/v1")
	api.Use(CommonSet)

	api.POST("/set/desc", v1.PostCardSets)
	api.GET("/card/desc", v1.GetCardsDesc)
	api.POST("/card/desc", v1.PostCardsDesc)
	api.GET("/card/price", v1.GetCardsPrice)
	api.POST("/card/price", v1.PostCardsPrice)
	api.POST("/card/pricewithimg", v1.PostCardsPriceWithDtcgDBImg)

	api.POST("/login", v1.Login)

	auth := api.Group("", middlewares.Auth)
	auth.GET("/me", v1.CurrentUser)
	auth.GET("/users/info/", v1.ListUser)

	auth.POST("/deck/price/json", v1.PostDeckPriceWithJSON)
	auth.GET("/deck/price/hid/:hid", v1.GetDeckPriceWithHID)
	auth.GET("/deck/price/cdid/:cdid", v1.GetDeckPriceWithCloudDeckID)
	auth.GET("/deck/price/wlid/:wlid", v1.GetDeckPriceWithJHSWishListID)
	auth.GET("/deck/price/share/:shareid", v1.GetDeckPriceWithShareID)
	auth.POST("/deck/price/ids", v1.PostDeckPriceWithIDS) // 根据 /deck/converter/:hid 接口转换后的字符串格式的卡组信息，获取卡组价格。

	// 将 DTCG DB 卡组广场中卡组的 URL 中最后的 HID，转变为由 card_id_from_db 组成的纯字符串格式。
	auth.GET("/deck/converter/:hid", v1.GetDeckConverter)

	auth.GET("/user/info/:uid", v1.GetUser)

	jhsAPI := auth.Group("/jhs")
	jhsAPI.GET("/market/sellers/products", v1.SellersProductsList)
	jhsAPI.PUT("/market/sellers/products/:productid", v1.SellersProductsUpdate)

	return r
}
