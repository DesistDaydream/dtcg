package models

type ProductsAddRequestBody struct {
	CardVersionID        string `query:"card_version_id"`
	Price                string `query:"price"`
	Quantity             string `query:"quantity"`
	Condition            string `query:"condition"`
	Remark               string `query:"remark"`
	GameKey              string `query:"game_key"`
	UserCardVersionImage string `query:"user_card_version_image"`
}
