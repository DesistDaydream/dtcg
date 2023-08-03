package models

// 获取推荐列表的请求参数
type WishListRecommendReqQuery struct {
	GameKey     string `form:"game_key"`
	GameSubKey  string `form:"game_sub_key"`
	IsRecommend string `form:"is_recommend"`
	Page        string `form:"page"`
}
