package models

type CardFilterInfoReq struct {
	// DTCG 编号为 2
	GameTitleID string `json:"game_title_id"`
	// 默认 EN
	LanguageCode string `json:"language_code"`
}
