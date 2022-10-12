package database

import "github.com/sirupsen/logrus"

type CardsDescFromDtcgDB struct {
	Count       int64 `json:"count"`
	PageSize    int
	PageCurrent int
	PageTotal   int
	Data        []CardDescFromDtcgDB `json:"data"`
}

type CardDescFromDtcgDB struct {
	ID             int    `gorm:"primaryKey" json:"my_id"` // ID
	CardID         int    `json:"card_id"`
	CardPack       int    `json:"card_pack"`
	PackName       string `json:"pack_name"`
	PackPrefix     string `json:"pack_prefix"`
	Serial         string `json:"serial"`
	SubSerial      string `json:"sub_serial"`
	JapName        string `json:"jap_name"`
	ScName         string `json:"sc_name"`
	Rarity         string `json:"rarity"`
	Type           string `json:"type"`
	Color          string `json:"color"`
	Level          string `json:"level"`
	Cost           string `json:"cost"`
	Cost1          string `json:"cost_1"`
	EvoCond        string `json:"evo_cond"`
	DP             string `json:"dp"`
	Grade          string `json:"grade"`
	Attribute      string `json:"attribute"`
	Class          string `json:"class"`
	Illustrator    string `json:"illustrator"`
	Effect         string `json:"effect"`
	EvoCoverEffect string `json:"evo_cover_effect"`
	SecurityEffect string `json:"security_effect"`
	IncludeInfo    string `json:"include_info"`
	RaritySC       string `json:"rarity_sc"`
}

func AddCardDescFromDtcgDB(cardDescFromDtcgDB *CardDescFromDtcgDB) {
	result := db.FirstOrCreate(cardDescFromDtcgDB, cardDescFromDtcgDB)
	if result.Error != nil {
		logrus.Errorf("插入数据失败: %v", result.Error)
	}
}

// 获取所有卡片描述
func ListCardDescFromDtcgDB() (*CardsDescFromDtcgDB, error) {
	var cd []CardDescFromDtcgDB
	result := db.Find(&cd)
	if result.Error != nil {
		return nil, result.Error
	}

	return &CardsDescFromDtcgDB{
		Count: result.RowsAffected,
		Data:  cd,
	}, nil
}

// 根据条件获取卡片描述
func GetCardDescFromDtcgDB(pageSize int, pageNum int) (*CardsDescFromDtcgDB, error) {
	var (
		CardCount int64
		cd        []CardDescFromDtcgDB
	)

	db.Model(&CardDescFromDtcgDB{}).Count(&CardCount)

	result := db.Limit(pageSize).Offset(pageSize * (pageNum - 1)).Find(&cd)
	if result.Error != nil {
		return nil, result.Error
	}

	return &CardsDescFromDtcgDB{
		Count:       CardCount,
		PageSize:    pageSize,
		PageCurrent: pageNum,
		PageTotal:   (int(CardCount) / pageSize) + 1,
		Data:        cd,
	}, nil
}
