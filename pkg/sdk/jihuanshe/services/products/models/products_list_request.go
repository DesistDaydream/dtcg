package models

type ProductsGetRequestQuery struct {
	GameKey       string `game_key`
	SellerUserID  string `seller_user_id`
	CardVersionID string `card_version_id`
	Token         string `token`
}
