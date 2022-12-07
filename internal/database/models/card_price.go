package models

import "time"

type CardsPrice struct {
	Count       int64       `json:"count"`
	PageSize    int         `json:"page_size"`
	PageCurrent int         `json:"page_current"`
	PageTotal   int         `json:"page_total"`
	Data        []CardPrice `json:"data"`
}

type CardPrice struct {
	ID             int       `gorm:"primaryKey" json:"id"`
	CardIDFromDB   int       `json:"card_id_from_db"`
	SetID          int       `json:"set_id"`
	SetPrefix      string    `json:"set_prefix"`
	Serial         string    `json:"serial"`
	ScName         string    `json:"sc_name"`
	AlternativeArt string    `json:"alternative_art"`
	Rarity         string    `json:"rarity"`
	CardVersionID  int       `json:"card_version_id"`
	MinPrice       float64   `json:"min_price"`
	AvgPrice       float64   `json:"avg_price"`
	CreatedAt      time.Time `json:"create_at"`
	UpdatedAt      time.Time `json:"update_at"`
	ImageUrl       string    `json:"image_url"`
}

// 查询卡牌价格的条件
type QueryCardPrice struct {
	CardPack       int64     `json:"card_pack"`
	ClassInput     bool      `json:"class_input"`
	Color          []string  `json:"color"`
	EvoCond        []EvoCond `json:"evo_cond"`
	Keyword        string    `json:"keyword"`
	Language       string    `json:"language"`
	OrderType      string    `json:"order_type"`
	QField         []string  `json:"qField"` // 通过 Keyword 进行查询的字段
	Rarity         []string  `json:"rarity"`
	Tags           []string  `json:"tags"` // 特征
	TagsLogic      string    `json:"tags__logic"`
	Type           string    `json:"type"`
	AlternativeArt string    `json:"alternative_art"`
}

// 带有dtcg数据库中图片的卡牌价格
type CardsPriceWithImageDB struct {
	Count       int64                  `json:"count"`
	PageSize    int                    `json:"page_size"`
	PageCurrent int                    `json:"page_current"`
	PageTotal   int                    `json:"page_total"`
	Data        []CardPriceWithImageDB `json:"data"`
}

type CardPriceWithImageDB struct {
	// ID             int       `gorm:"primaryKey" json:"id"`
	CardIDFromDB int `json:"card_id_from_db"`
	// SetID          int       `json:"set_id"`
	SetPrefix      string  `json:"set_prefix"`
	Serial         string  `json:"serial"`
	ScName         string  `json:"sc_name"`
	AlternativeArt string  `json:"alternative_art"`
	Rarity         string  `json:"rarity"`
	CardVersionID  int     `json:"card_version_id"`
	MinPrice       float64 `json:"min_price"`
	AvgPrice       float64 `json:"avg_price"`
	// CreatedAt      time.Time `json:"create_at"`
	// UpdatedAt      time.Time `json:"update_at"`
	// ImageUrl       string    `json:"image_url"`
	Image string `json:"image"`
}
