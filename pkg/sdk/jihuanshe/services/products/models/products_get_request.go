package models

type ProductsListRequestQuery struct {
	GameKey    string `game_key`
	GameSubKey string `game_sub_key`
	OnSale     string `on_sale`
	Page       string `page`
	Token      string `token`
}
