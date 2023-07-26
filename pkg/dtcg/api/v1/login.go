package v1

import (
	"github.com/DesistDaydream/dtcg/pkg/dtcg/utils"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	token, err := utils.GenerateToken("DesistDaydream")
	if err != nil {
		utils.ErrorResp(c, err, 400, true)
		return
	}
	c.JSON(200, gin.H{"token": token})
}
