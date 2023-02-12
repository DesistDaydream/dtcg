package models

type ProductsGetReqQuery struct {
	SellerUserID  string `query:"seller_user_id"`
	CardVersionID string `query:"card_version_id"`
	GameKey       string `query:"game_key"`
}
