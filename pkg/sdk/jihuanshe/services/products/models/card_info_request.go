package models

type ProductsGetReqQuery struct {
	GameKey       string `query:"game_key"`
	SellerUserID  string `query:"seller_user_id"`
	CardVersionID string `query:"card_version_id"`
	Token         string `query:"token"`
}
