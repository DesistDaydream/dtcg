package models

import "time"

// 数据库模型。卡牌价格信息
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

// 卡牌价格的查询条件
type CardPriceQuery struct {
	CardVersionID      int      `json:"card_version_id"`        // 卡牌在集换社中的 ID
	NotInCardVersionID []int    `json:"not_in_card_version_id"` // 不包含
	SetsPrefix         []string `json:"set_prefix"`             // 卡牌集合
	Keyword            string   `json:"keyword"`                // 关键字
	Language           string   `json:"language"`               // 语言
	QField             []string `json:"qField"`                 // 通过 Keyword 进行查询的字段
	Rarity             []string `json:"rarity"`                 // 卡牌稀有度
	AlternativeArt     string   `json:"alternative_art"`        // 是否是异画。可用的值有两个：是、否
	MinPriceRange      string   `json:"min_price_range"`        // 最低价范围
	AvgPriceRange      string   `json:"avg_price_range"`        // 集换价范围
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
	CardIDFromDB   int     `json:"card_id_from_db"`
	SetPrefix      string  `json:"set_prefix"`
	Serial         string  `json:"serial"`
	ScName         string  `json:"sc_name"`
	AlternativeArt string  `json:"alternative_art"`
	Rarity         string  `json:"rarity"`
	CardVersionID  int     `json:"card_version_id"`
	MinPrice       float64 `json:"min_price"`
	AvgPrice       float64 `json:"avg_price"`
	// ImageUrl       string    `json:"image_url"`
	Image string `json:"image"`
}
