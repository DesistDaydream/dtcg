package v1

import (
	"net/http"
	"strconv"

	dbmodels "github.com/DesistDaydream/dtcg/internal/database/models"

	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/pkg/dtcg/api/v1/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 列出所有卡牌价格详情，分页
func GetCardsPrice(c *gin.Context) {
	// 绑定 url query
	// var reqQuery models.CommonReqQuery

	// if err := c.ShouldBindQuery(&reqQuery); err != nil {
	// 	logrus.Error(err)
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	resp, err := database.GetCardsPriceWithPaginationLib(c)
	if err != nil {
		logrus.Errorf("%v", err)
	}

	c.JSON(200, resp)
}

// 根据条件列出卡牌价格详情，分页
func PostCardsPrice(c *gin.Context) {
	pageNum, _ := strconv.Atoi(c.DefaultQuery("page_num", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	// 绑定请求体
	var reqBody dbmodels.CardPriceQuery
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ReqBodyErrorResp{
			Message: "请求体错误",
			Data:    "",
		})
		return
	}

	resp, err := database.GetCardPriceByCondition(pageSize, pageNum, &reqBody)
	if err != nil {
		logrus.Errorf("%v", err)
	}

	c.JSON(200, resp)
}

// 根据条件获取带有数码宝贝数据库中卡图的卡牌价格数据
func PostCardsPriceWithDtcgDBImg(c *gin.Context) {
	// 绑定 url query
	var reqQuery models.CommonReqQuery
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
	var reqBody dbmodels.CardPriceQuery
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
