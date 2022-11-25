package v1

import (
	"net/http"

	dbmodels "github.com/DesistDaydream/dtcg/internal/database/models"

	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/pkg/dtcg/api/v1/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 获取所有卡牌价格详情
func GetCardsPrice(c *gin.Context) {
	// 允许跨域
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	// 绑定 url query
	var req models.GetCardsPriceReqQuery
	if err := c.ShouldBindQuery(&req); err != nil {
		logrus.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.PageSize == 0 || req.PageNum == 0 {
		req.PageSize = 10
		req.PageNum = 1
	}

	resp, err := database.GetCardsPrice(req.PageSize, req.PageNum)
	if err != nil {
		logrus.Errorf("%v", err)
	}

	c.JSON(200, resp)
}

// 根据条件获取卡牌价格详情
func PostCardsPrice(c *gin.Context) {
	// 允许跨域
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	// 绑定 url query
	var reqQuery models.GetCardsPriceReqQuery
	if err := c.ShouldBindQuery(&reqQuery); err != nil {
		logrus.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if reqQuery.PageSize == 0 || reqQuery.PageNum == 0 {
		reqQuery.PageSize = 10
		reqQuery.PageNum = 1
	}

	// 绑定请求体
	var reqBody dbmodels.QueryCardPrice
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ReqBodyErrorResp{
			Message: "请求体错误",
			Data:    "",
		})
		return
	}

	resp, err := database.GetCardPriceByCondition(reqQuery.PageSize, reqQuery.PageNum, &reqBody)
	if err != nil {
		logrus.Errorf("%v", err)
	}

	c.JSON(200, resp)
}

// 根据条件获取带有数码宝贝数据库中卡图的卡牌价格数据
func PostCardsPriceWithDtcgDBImg(c *gin.Context) {
	// 允许跨域
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	// 绑定 url query
	var reqQuery models.GetCardsPriceReqQuery
	if err := c.ShouldBindQuery(&reqQuery); err != nil {
		logrus.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if reqQuery.PageSize == 0 || reqQuery.PageNum == 0 {
		reqQuery.PageSize = 10
		reqQuery.PageNum = 1
	}

	// 绑定请求体
	var reqBody dbmodels.QueryCardPrice
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ReqBodyErrorResp{
			Message: "请求体错误",
			Data:    "",
		})
		return
	}

	// 根据条件获取带有数码宝贝数据库中卡图的卡牌价格数据
	resp, err := database.GetCardPriceWithDtcgDBImgByCondition(reqQuery.PageSize, reqQuery.PageNum, &reqBody)
	if err != nil {
		logrus.Errorf("%v", err)
	}

	c.JSON(200, resp)
}
