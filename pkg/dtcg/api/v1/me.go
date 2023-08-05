package v1

import (
	"github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/gin-gonic/gin"
)

type UserResp struct {
	models.User
	Otp bool `json:"otp"`
}

func CurrentUser(c *gin.Context) {
	user := c.MustGet("user").(*models.User)
	userResp := UserResp{
		User: *user,
	}
	userResp.Password = ""

	c.JSON(200, userResp)
}
