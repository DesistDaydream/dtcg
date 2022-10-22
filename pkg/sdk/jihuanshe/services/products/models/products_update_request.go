package models

type ProductsUpdateReqBody struct {
	Condition            string `json:"condition"`
	OnSale               string `json:"on_sale"`
	Price                string `json:"price"`
	Quantity             string `json:"quantity"`
	Remark               string `json:"remark"`
	UserCardVersionImage string `json:"user_card_version_image"`
}

type ProductsUpdateReqQuery struct {
	Token string `query:"token"`
}
