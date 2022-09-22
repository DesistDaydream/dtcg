package models

type ProductsUpdateRequestBody struct {
	Condition            string `query:"condition"`
	OnSale               string `query:"on_sale"`
	Price                string `query:"price"`
	Quantity             string `query:"quantity"`
	Remark               string `query:"remark"`
	UserCardVersionImage string `query:"user_card_version_image"`
}

type ProductsUpdateRequestQuery struct {
	Token string `query:"token"`
}
