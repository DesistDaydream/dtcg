package v1

import (
	"fmt"
	"net/http"

	"github.com/DesistDaydream/dtcg/pkg/dtcg/api/v1/models"
	"github.com/DesistDaydream/dtcg/pkg/dtcg/deck"
	"github.com/gin-gonic/gin"
)

// 根据 DTCG DB 卡组广场中卡组的 URL 中最后的 HID，将卡组信息转变为由 card_id_from_db 组成的纯字符串格式。
// 简单点说，根据 HID 导出字符串格式的卡组信息。其实也可以说是 JSON 格式的。官方导出的是卡牌编号，我导出的是卡牌 ID，这样可以避免异画卡无法识别的问题。
func GetDeckConverter(c *gin.Context) {
	// 允许跨域
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	// 获取 URL 中的 HID
	hid := c.Param("hid")

	resp, err := deck.GetResp(hid)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ReqBodyErrorReponse{
			Message: "获取响应失败",
			Data:    fmt.Sprintf("%v", err),
		})
		return
	}

	c.JSON(200, &resp)
}
