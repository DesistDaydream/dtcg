package models

type ProductsGetReqQuery struct {
	SellerUserID  string `form:"seller_user_id"`
	CardVersionID string `form:"card_version_id"`
	GameKey       string `form:"game_key"`
}
