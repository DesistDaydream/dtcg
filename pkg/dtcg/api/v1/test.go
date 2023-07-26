package v1

import (
	"github.com/gin-gonic/gin"
)

func TestAuth(c *gin.Context) {
	c.JSON(200, "认证通过")
}
