package models

import "time"

type CardMetadataGetResp struct {
	Success CardMetadataGetRespData `json:"success"`
}

type CardMetadataGetRespData struct {
	CardSearchConfigs []CardSearchConfig `json:"card_search_configs"`
	CardSetList       []CardSetList      `json:"card_set_list"`
	Code              int64              `json:"code"`
	GameFormats       []GameFormat       `json:"game_formats"`
	IsDivisionText    bool               `json:"is_division_text"`
	IsReverseCard     bool               `json:"is_reverse_card"`
}

// 卡牌搜索所需信息
type CardSearchConfig struct {
	Choices        []string `json:"choices"`
	ConfigName     string   `json:"config_name"`
	ConfigNumber   string   `json:"config_number"`
	DisplayColumns int64    `json:"display_columns"`
	SearchType     int64    `json:"search_type"`
}

// 卡集列表
type CardSetList struct {
	CreatedAt    time.Time `json:"created_at"`
	Explain      string    `json:"explain"`
	Flgs         int64     `json:"flgs"`
	GameTitleID  int64     `json:"game_title_id"`
	ID           int64     `json:"id"`
	ImageURL     string    `json:"image_url"`
	LanguageCode string    `json:"language_code"`
	Name         string    `json:"name"`
	Number       string    `json:"number"`
	ReleaseDate  time.Time `json:"release_date"`
	SortOrder    int64     `json:"sort_order"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type GameFormat struct {
	FormatName   string `json:"format_name"`
	GameFormatID int64  `json:"game_format_id"`
}
