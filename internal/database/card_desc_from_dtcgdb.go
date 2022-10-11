package database

import "github.com/sirupsen/logrus"

type CardsDescFromDtcgDB struct {
	Data    CardsDescFromDtcgDBData `json:"data"`
	Message string                  `json:"message"`
	Success bool                    `json:"success"`
}

type CardsDescFromDtcgDBData struct {
	Count int64                `json:"count"`
	List  []CardDescFromDtcgDB `json:"list"`
}

type CardDescFromDtcgDB struct {
	Attribute      string   `json:"attribute"`
	CardID         int64    `json:"card_id"`
	CardPack       int64    `json:"card_pack"`
	Class          []string `json:"class" gorm:"type:text"`
	Color          []string `json:"color" gorm:"type:text"`
	Cost           string   `json:"cost"`
	Cost1          string   `json:"cost_1"`
	DP             string   `json:"DP"`
	Effect         string   `json:"effect"`
	EvoCond        string   `json:"evo_cond"`
	EvoCoverEffect string   `json:"evo_cover_effect"`
	Grade          string   `json:"grade"`
	Illustrator    string   `json:"illustrator"`
	// Images         []Image  `json:"images" gorm:"type:text"`
	IncludeInfo string `json:"include_info"`
	JapName     string `json:"japName"`
	Level       string `json:"level"`
	// Package        Package `json:"package" gorm:"type:text"`
	Rarity         string `json:"rarity"`
	RaritySC       string `json:"rarity$SC"`
	ScName         string `json:"scName"`
	SecurityEffect string `json:"security_effect"`
	Serial         string `json:"serial"`
	SubSerial      string `json:"sub_serial"`
	Type           string `json:"type"`
}

type Image struct {
	CardID    int64  `json:"card_id"`
	ID        int64  `json:"id"`
	ImgPath   string `json:"img_path"`
	ImgRare   string `json:"img_rare"`
	ThumbPath string `json:"thumb_path"`
}

type Package struct {
	Language   string `json:"language"`
	PackID     int64  `json:"pack_id"`
	PackPrefix string `json:"pack_prefix"`
}

func AddCardDescFromDtcgDB(cardDescFromDtcgDB *CardDescFromDtcgDB) {
	result := db.FirstOrCreate(cardDescFromDtcgDB, cardDescFromDtcgDB)
	if result.Error != nil {
		logrus.Errorf("插入数据失败: %v", result.Error)
	}
}
