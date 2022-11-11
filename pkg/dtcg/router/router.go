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
	rr.POST("/deck/price", v1.PostDeckPrice)
	rr.POST("/deck/price/ids", v1.PostDeckPriceWithID)
	rr.GET("/deck/price/hid/:hid", v1.GetDeckPriceWithHID)
	rr.GET("/deck/converter/:hid", v1.GetDeckConverter)

	return r
}
