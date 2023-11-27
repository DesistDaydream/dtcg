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

// 列出所有卡牌的描述，分页
func GetCardsDesc(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	pageNum, _ := strconv.Atoi(c.DefaultQuery("page_num", "1"))

	resp, err := database.GetCardsDesc(pageSize, pageNum)
	if err != nil {
		logrus.Errorf("%v", err)
	}

	c.JSON(200, resp)
}

// 根据条件获取卡牌描述详情
func PostCardsDesc(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	pageNum, _ := strconv.Atoi(c.DefaultQuery("page_num", "1"))

	// 绑定请求体
	var reqBody dbmodels.CardDescQuery
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		logrus.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ReqBodyErrorResp{
			Message: "请求体错误",
			Data:    "",
		})
		return
	}

	resp, err := database.GetCardDescByCondition(pageSize, pageNum, &reqBody)
	if err != nil {
		logrus.Errorf("%v", err)
	}

	c.JSON(200, resp)
}
