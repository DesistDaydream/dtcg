package models

type ProductsListReqQuery struct {
	GameKey    string `form:"game_key"`
	GameSubKey string `form:"game_sub_key"`
	Keyword    string `form:"keyword"`
	OnSale     string `form:"on_sale"` // 售卖状态。1: 在售，0: 下架
	Page       string `form:"page"`
	Sorting    string `form:"sorting"` // 排序逻辑。published_at_desc,price_desc,price_asc。默认值: published_at_desc。published_at_desc 是按照上架时间从新到旧，其他值是按照上架时间从旧到新。
	Rarity     string `form:"rarity"`  // 商品罕贵度。可用的值有：异画,特典,SEC,SR,R,U,C,P。默认值: ""。
}
