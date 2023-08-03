package v1

import (
	"github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func TestAuth(c *gin.Context) {
	user := c.MustGet("user").(*models.User)

	logrus.Infof("%v", user)

	c.JSON(200, "认证通过")
}
