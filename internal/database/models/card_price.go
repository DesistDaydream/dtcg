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
}