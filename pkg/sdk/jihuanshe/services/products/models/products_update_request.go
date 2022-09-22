package models

type ProductsUpdateRequestBody struct {
	Condition            int64  `query:"condition"`
	OnSale               int64  `query:"on_sale"`
	Price                string `query:"price"`
	Quantity             int64  `query:"quantity"`
	Remark               string `query:"remark"`
	UserCardVersionImage string `query:"user_card_version_image"`
}

type ProductsUpdateRequestQuery struct {
	Token string `query:"token"`
}
