package models

type ProductsListReqQuery struct {
	GameKey    string `query:"game_key"`
	GameSubKey string `query:"game_sub_key"`
	Keyword    string `query:"keyword"`
	OnSale     string `query:"on_sale"`
	Page       string `query:"page"`
	Sorting    string `query:"sorting"`
	Token      string `query:"token"`
}
