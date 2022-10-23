package models

type ProductsAddReqBody struct {
	CardVersionID        string `json:"card_version_id"`
	Price                string `json:"price"`
	Quantity             string `json:"quantity"`
	Condition            string `json:"condition"`
	Remark               string `json:"remark"`
	GameKey              string `json:"game_key"`
	UserCardVersionImage string `json:"user_card_version_image"`
}

type ProductsAddReqQuery struct {
	Token string `query:"token"`
}
