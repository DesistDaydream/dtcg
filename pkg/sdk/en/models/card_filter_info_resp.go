package models

import "time"

type CardFilterInfo struct {
	Success CardFilterInfoData `json:"success"`
}

type CardFilterInfoData struct {
	Code              int                 `json:"code"`
	GameFormats       []GameFormats       `json:"game_formats"`
	CardSearchConfigs []CardSearchConfigs `json:"card_search_configs"`
	CardSetList       []CardSetList       `json:"card_set_list"`
}

type GameFormats struct {
	GameFormatID int    `json:"game_format_id"`
	FormatName   string `json:"format_name"`
}

// 卡片过滤配置
type CardSearchConfigs struct {
	SearchType   int      `json:"search_type"`
	ConfigName   string   `json:"config_name"`
	ConfigNumber string   `json:"config_number"`
	Choices      []string `json:"choices,omitempty"`
}

// 卡包列表
type CardSetList struct {
	GameTitleID  int       `json:"game_title_id"`
	Number       string    `json:"number"`
	Name         string    `json:"name"`
	Explain      string    `json:"explain"`
	ImageURL     string    `json:"image_url"`
	LanguageCode string    `json:"language_code"`
	ReleaseDate  time.Time `json:"release_date"`
	SortOrder    int       `json:"sort_order"`
	ID           int       `json:"id"`
	Flgs         int       `json:"flgs"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
