package models

// 获取推荐列表的请求参数
type WishListRecommendReqQuery struct {
	GameKey     string `query:"game_key"`
	GameSubKey  string `query:"game_sub_key"`
	IsRecommend string `query:"is_recommend"`
	Page        string `query:"page"`
}
