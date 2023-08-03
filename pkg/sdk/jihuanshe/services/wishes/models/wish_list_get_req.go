package models

// 获取推荐列表的请求参数
type WishListGetReqQuery struct {
	GameKey    string `form:"game_key"`
	GameSubKey string `form:"game_sub_key"`
	Page       string `form:"page"`
	WishListID string `form:"wish_list_id"`
}
