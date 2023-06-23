package models

// 获取推荐列表的请求参数
type WishListGetReqQuery struct {
	GameKey    string `query:"game_key"`
	GameSubKey string `query:"game_sub_key"`
	Page       string `query:"page"`
	WishListID string `query:"wish_list_id"`
}
