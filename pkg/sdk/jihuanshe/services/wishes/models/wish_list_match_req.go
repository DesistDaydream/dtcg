package models

// 获取推荐列表的请求参数
type WishListMatchResultsReqQuery struct {
	GameKey           string `query:"game_key"`
	GameSubKey        string `query:"game_sub_key"`
	IgnoreCardVersion string `query:"ignore_card_version"`
	ShowMatchDetails  string `query:"show_match_details"` // 是否显示匹配到的细节，即显示匹配到的每张卡牌的信息。1: 显示；0: 不显示
	WishListID        string `query:"wish_list_id"`
}
