package models

type ProductSellersGetReqQuery struct {
	CardVersionID string `query:"card_version_id"`
	Condition     string `query:"condition"` // 1 为在售，0 为下架
	GameKey       string `query:"game_key"`
	Page          string `query:"page"`
}
