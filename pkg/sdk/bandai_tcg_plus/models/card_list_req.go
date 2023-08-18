package models

type CardListReqQuery struct {
	CardSet     string `form:"card_set[]"`    // 卡包编号
	GameTitleID string `form:"game_title_id"` // DTCG 编号。
	Limit       string `form:"limit"`         // 每次返回卡片的数量。1-1000 的整数。
	Offset      string `form:"offset"`        // 偏移量。每次返回的卡片从偏移量开始。实现了类似于分页的效果
}
