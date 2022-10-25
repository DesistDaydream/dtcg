package models

type ProductSellersGetReqQuery struct {
	CardVersionID string `query:"card_version_id"`
	// 1 为在售，0 为下架
	Condition string `query:"condition"`
	GameKey   string `query:"game_key"`
	Page      string `query:"page"`
}
