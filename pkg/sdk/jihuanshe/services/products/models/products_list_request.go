package models

type ProductsListReqQuery struct {
	GameKey    string `query:"game_key"`
	GameSubKey string `query:"game_sub_key"`
	Keyword    string `query:"keyword"`
	OnSale     string `query:"on_sale"` // 售卖状态。1: 在售，0: 下架
	Page       string `query:"page"`
	Sorting    string `query:"sorting"`
}
