package models

type ProductSellersGetReqQuery struct {
	CardVersionID string `form:"card_version_id"`
	Condition     string `form:"condition"` // 1 为在售，0 为下架
	GameKey       string `form:"game_key"`
	Page          string `form:"page"`
}
