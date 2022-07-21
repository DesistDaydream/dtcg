package models

type CardListReq struct {
	// 卡包编号
	CardSet string `json:"card_set"`
	// DTCG 编号为 2
	GameTitleID string
	// 每次返回卡片的数量。1-1000 的整数。
	Limit string
	// 偏移量。每次返回的卡片从偏移量开始。
	Offset string
}
