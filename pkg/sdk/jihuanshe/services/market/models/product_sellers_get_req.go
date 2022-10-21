package models

type ProductSellersGetReqQuery struct {
	CardVersionID string `query:"card_version_id"`
	Condition     string `query:"condition"`
	GameKey       string `query:"game_key"`
	Page          string `query:"page"`
}
