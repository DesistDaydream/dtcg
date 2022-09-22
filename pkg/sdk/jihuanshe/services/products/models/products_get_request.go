package models

type ProductsListRequestQuery struct {
	GameKey    string `query:"game_key"`
	GameSubKey string `query:"game_sub_key"`
	OnSale     string `query:"on_sale"`
	Page       string `query:"page"`
	Token      string `query:"token"`
}
