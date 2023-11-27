package v1

import (
	"strconv"

	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 获取所有卡牌集合的信息
func GetCardSets(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	pageNum, _ := strconv.Atoi(c.DefaultQuery("page_num", "1"))

	resp, err := database.GetCardSets(pageSize, pageNum)
	if err != nil {
		logrus.Errorf("%v", err)
	}

	c.JSON(200, resp)
}
