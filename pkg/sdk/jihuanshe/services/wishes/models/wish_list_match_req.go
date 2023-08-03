package models

// 获取推荐列表的请求参数
type WishListMatchResultsReqQuery struct {
	GameKey           string `form:"game_key"`
	GameSubKey        string `form:"game_sub_key"`
	IgnoreCardVersion string `form:"ignore_card_version"`
	ShowMatchDetails  string `form:"show_match_details"` // 是否显示匹配到的细节，即显示匹配到的每张卡牌的信息。1: 显示；0: 不显示
	WishListID        string `form:"wish_list_id"`
}
