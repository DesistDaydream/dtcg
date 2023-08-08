package models

// 列出卡牌
type CardVersionsListResp struct {
	CurrentPage int64                     `json:"current_page"`
	Data        []CardVersionListRespData `json:"data"`
	From        int64                     `json:"from"`
	LastPage    int64                     `json:"last_page"`
	NextPageURL string                    `json:"next_page_url"`
	Path        string                    `json:"path"`
	PerPage     int64                     `json:"per_page"`
	PrevPageURL string                    `json:"prev_page_url"`
	To          int64                     `json:"to"`
	Total       int64                     `json:"total"`
}

type CardVersionListRespData struct {
	AvgPrice      string     `json:"avg_price"`
	CardID        int64      `json:"card_id"`
	CardNames     []CardName `json:"card_names"`
	CardVersionID int64      `json:"card_version_id"`
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
