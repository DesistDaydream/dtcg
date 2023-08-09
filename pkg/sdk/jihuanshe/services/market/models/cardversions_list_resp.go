package models

// 列出卡牌
type CardVersionsListResp struct {
	CurrentPage int64                     `json:"current_page"`
	Data        []CardVersionListRespData `json:"data"`
	From        int                       `json:"from"`
	LastPage    int                       `json:"last_page"`
	NextPageURL string                    `json:"next_page_url"`
	Path        string                    `json:"path"`
	PerPage     int                       `json:"per_page"`
	PrevPageURL string                    `json:"prev_page_url"`
	To          int                       `json:"to"`
	Total       int                       `json:"total"`
}

// 注意 CardID 和 CardVersionID 的区别
type CardVersionListRespData struct {
	AvgPrice      string     `json:"avg_price"`
	CardID        int        `json:"card_id"` // 卡牌 ID。一张卡牌的原画和异画都是同一个 卡牌ID
	CardNames     []CardName `json:"card_names"`
	CardVersionID int        `json:"card_version_id"` // 卡牌版本 ID。全局唯一标识符，同一张卡牌的原画和异画，其卡牌版本ID不同
	Grade         string     `json:"grade"`
	ImageURL      string     `json:"image_url"`
	Language      string     `json:"language"`
	LanguageText  string     `json:"language_text"`
	MinPrice      string     `json:"min_price"`
	NameCN        string     `json:"name_cn"`
	NameOrigin    string     `json:"name_origin"`
	Number        string     `json:"number"`
	NumberAlias   string     `json:"number_alias"`
	Rarity        string     `json:"rarity"`
}
