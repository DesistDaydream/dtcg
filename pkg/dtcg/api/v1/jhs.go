package v1

import (
	"fmt"
	"net/http"

	"github.com/DesistDaydream/dtcg/pkg/handler"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/market/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 列出我在卖的商品
func SellersProductsList(c *gin.Context) {
	fmt.Println(c.Request.URL.Query())
	var req models.ProductsListReqQuery
	if err := c.ShouldBindQuery(&req); err != nil {
		logrus.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("啦啦啦", req.Page, req.Keyword, req.OnSale, req.Sorting)

	resp, err := handler.H.JhsServices.Market.SellersProductsList(req.Page, req.Keyword, req.OnSale, req.Sorting)
	if err != nil {
		logrus.Errorf("%v", err)
	}

	c.JSON(200, resp)
}

// 更新我在卖的商品
func SellersProductsUpdate(c *gin.Context) {
	productid := c.Param("productid")

	var req models.ProductsUpdateReqBody
	if err := c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("绑定请求体异常，原因: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := handler.H.JhsServices.Market.SellersProductsUpdate(&req, productid)
	if err != nil {
		logrus.Errorf("%v", err)
	}

	c.JSON(200, resp)
}
