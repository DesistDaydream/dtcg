package models

type DeckSearchReqQuery struct {
	Limit string `form:"limit"`
	Page  string `form:"page"`
}

type DeckSearchReqBody struct {
	Envir string   `json:"envir"` // 环境。chs: 简中；ja: 日文
	Kw    string   `json:"kw"`    // 搜索关键字
	Tags  []string `json:"tags"`  // 共两个元素，分别是类别与颜色。3: 卡组分享；13: 国内比赛。5: 红色；不填: 无限制。
}
