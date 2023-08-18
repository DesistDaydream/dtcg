package models

type CardFilterInfoReq struct {
	// DTCG 编号
	GameTitleID string `json:"game_title_id"` // 2: DTCG 英文; 6: DTCG 日文
	// 默认 EN
	LanguageCode string `json:"language_code"`
}

type CardMetadataGetReqQuery struct {
	GameTitleID  string `form:"game_title_id"` // DTCG 编号。2: DTCG 英文; 6: DTCG 日文
	LanguageCode string `form:"language_code"` // 默认 EN
}
