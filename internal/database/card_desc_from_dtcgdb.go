package database

import "github.com/sirupsen/logrus"

type CardsDescFromDtcgDB struct {
	Count int64                `json:"count"`
	List  []CardDescFromDtcgDB `json:"list"`
}

type CardDescFromDtcgDB struct {
	ID             int    `gorm:"primaryKey" json:"my_id"` // ID
	CardID         int    `json:"card_id"`
	CardPack       int    `json:"card_pack"`
	Serial         string `json:"serial"`
	SubSerial      string `json:"sub_serial"`
	JapName        string `json:"japName"`
	ScName         string `json:"scName"`
	Rarity         string `json:"rarity"`
	Type           string `json:"type"`
	Color          string `json:"color"`
	Level          string `json:"level"`
	Cost           string `json:"cost"`
	Cost1          string `json:"cost_1"`
	EvoCond        string `json:"evo_cond"`
	DP             string `json:"DP"`
	Grade          string `json:"grade"`
	Attribute      string `json:"attribute"`
	Class          string `json:"class"`
	Illustrator    string `json:"illustrator"`
	Effect         string `json:"effect"`
	EvoCoverEffect string `json:"evo_cover_effect"`
	SecurityEffect string `json:"security_effect"`
	IncludeInfo    string `json:"include_info"`
	RaritySC       string `json:"rarity$SC"`
}

func AddCardDescFromDtcgDB(cardDescFromDtcgDB *CardDescFromDtcgDB) {
	result := db.FirstOrCreate(cardDescFromDtcgDB, cardDescFromDtcgDB)
	if result.Error != nil {
		logrus.Errorf("插入数据失败: %v", result.Error)
	}
}
