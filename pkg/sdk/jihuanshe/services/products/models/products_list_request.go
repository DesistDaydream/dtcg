package models

type ProductsListReqQuery struct {
	GameKey    string `query:"game_key"`
	GameSubKey string `query:"game_sub_key"`
	Keyword    string `query:"keyword"`
	OnSale     string `query:"on_sale"` // 是否在售。1: 在售，0或其他数字: 下架
	Page       string `query:"page"`
	Sorting    string `query:"sorting"`
	// Token      string `query:"token"`
}
