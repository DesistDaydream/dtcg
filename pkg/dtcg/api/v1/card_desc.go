package v1

import (
	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetCardsDesc(c *gin.Context) {
	cardsDesc, err := database.ListCardDescFromDtcgDB()
	if err != nil {
		logrus.Errorf("%v", err)
	}

	c.JSON(200, cardsDesc)
}
