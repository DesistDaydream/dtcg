package models

type ProductsListReqQuery struct {
	GameKey    string `query:"game_key"`
	GameSubKey string `query:"game_sub_key"`
	Keyword    string `query:"keyword"`
	OnSale     string `query:"on_sale"` // 售卖状态。1: 在售，0: 下架
	Page       string `query:"page"`
	Sorting    string `query:"sorting"` // 可用的值有：price_desc,price_asc,空。空值是按照上架顺序排列
	Rarity     string `query:"rarity"`  // 商品的罕贵度。可用的值有：异画,特典,SEC,SR,R,U,C,P
}
